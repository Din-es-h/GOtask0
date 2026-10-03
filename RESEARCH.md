# Migration Research: mgo.v2 → mongo-driver/v2/mongo

This document compares the legacy `mgo.v2` driver used in `legacy-mgo` against the
official modern `go.mongodb.org/mongo-driver` used in `feature/modern-driver`,
covering syntax, architecture, and the reasoning behind what changed.

## 1. Syntax Comparison

| Operation            | mgo.v2 (legacy)                                              | Modern driver                                                        |
|-----------------------|---------------------------------------------------------------|------------------------------------------------------------------------|
| Connect               | `session, err := mgo.Dial("host")`                            | `client, err := mongo.Connect(ctx, options.Client().ApplyURI("uri"))` |
| Confirm reachable     | (done automatically inside `Dial`)                             | `client.Ping(ctx, nil)` — separate, explicit call                     |
| Get a collection      | `session.DB("name").C("coll")`                                 | `client.Database("name").Collection("coll")`                          |
| Insert one document   | `coll.Insert(doc)`                                              | `coll.InsertOne(ctx, doc)`                                             |
| Find one by ID        | `coll.FindId(id).One(&result)`                                  | `coll.FindOne(ctx, bson.M{"_id": id}).Decode(&result)`                |
| Delete one by ID      | `coll.RemoveId(id)`                                              | `coll.DeleteOne(ctx, bson.M{"_id": id})`                               |
| ID type               | `bson.ObjectId`                                                  | `primitive.ObjectID`                                                   |
| Generate new ID       | `bson.NewObjectId()`                                              | `primitive.NewObjectID()`                                              |
| Validate + convert ID from URL string | `bson.IsObjectIdHex(s)` then `bson.ObjectIdHex(s)` (two steps) | `primitive.ObjectIDFromHex(s)` (one call, returns an error if invalid) |
| Close/shut down       | `session.Close()`                                                 | `client.Disconnect(ctx)`                                               |

Every modern-driver operation above takes a `ctx context.Context` as its first
argument — mgo.v2 has no equivalent parameter anywhere.

## 2. Architectural Differences

**Session vs. Client.** mgo.v2 centres on a `Session`, dialed once and reused.
The modern driver replaces this with a `Client`, connected via `mongo.Connect`.
Functionally similar (one shared, long-lived handle reused across requests),
but the *behaviour* of connecting itself changed — see below.

**Synchronous dial vs. async connect + explicit ping.** `mgo.Dial` verifies the
server is reachable immediately, synchronously, before returning — if it can't
reach MongoDB, it fails right there. `mongo.Connect` does not necessarily
verify reachability at all; it starts background goroutines that monitor the
deployment's topology on an ongoing basis and returns without blocking on a
full handshake. `client.Ping(ctx, nil)` exists specifically to force a
one-off, synchronous "is this actually reachable right now" check, filling
the gap `Connect` deliberately leaves open.

This isn't an arbitrary API split. Production MongoDB deployments are often
replica sets or sharded clusters, where "which server do I talk to" can
change over time (failover, nodes joining/leaving). The driver's real job is
continuous background monitoring of that topology for the program's whole
lifetime, not a single one-time dial — so `Connect` reflects that ongoing
responsibility instead of pretending the job is a single, finished event.

**`context.Context` on every operation.** The modern driver requires a
`context.Context` as the first argument to essentially every database call.
A context carries an optional deadline/timeout and a cancellation signal.
Without it (as in mgo.v2), a stuck or overloaded server can leave a call
blocked indefinitely, with no way for the caller to give up early. With a
context, every operation can be bounded ("give up after 5 seconds") and tied
to the lifetime of whatever triggered it (e.g. cancel the DB call if the
original HTTP request was cancelled).

**Combined validation.** `bson.IsObjectIdHex` + `bson.ObjectIdHex` (check,
then convert) becomes a single `primitive.ObjectIDFromHex`, which returns an
error directly if the input isn't a valid ID — one call instead of two,
following the same `(result, error)` pattern used throughout the rest of
idiomatic Go.

**Naming convention fix.** `bson.ObjectId` (lowercase trailing letters)
becomes `primitive.ObjectID` (fully capitalised, per Go's own style
convention that initialisms like ID, URL, HTTP should stay fully uppercase
rather than being treated as an ordinary word). A small but genuine sign that
the rewrite cleaned up naming consistency, not just functionality.

## 3. Why the Change — Summary

mgo.v2 has been unmaintained since around 2018. The underlying MongoDB wire
protocol moved on since then: MongoDB deprecated the old `OP_QUERY` message
format in version 5.0 and removed server-side support for it entirely in
5.1. mgo.v2 was built around `OP_QUERY` and never updated — so it cannot
complete a handshake with any MongoDB server from 5.1 onward at all. This was
confirmed directly while building this project: `mgo.Dial` against a current
MongoDB server failed with "no reachable servers," even though the server was
verified running and listening (via `netstat`) the whole time — the failure
was a protocol-level handshake mismatch, not a networking problem.

Beyond protocol compatibility, the modern driver's design changes — explicit
contexts, async connect with separate ping, cluster-aware background
monitoring — reflect real production concerns (timeouts, replica sets,
graceful cancellation) that a single-server, one-shot `Dial` model didn't
address.
