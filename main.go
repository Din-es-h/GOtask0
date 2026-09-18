package main

import (
    "encoding/json"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func main() {
    var err error
    session, err = mgo.Dial("127.0.0.1:27018")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    router := httprouter.New()
    router.POST("/user", CreateUser)
    router.GET("/user/:id", GetUser)
    router.DELETE("/user/:id", DeleteUser)

    http.ListenAndServe(":8080", router)
}


type User struct {
    ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
    Name string        `bson:"name" json:"name"`
    Age  int           `bson:"age" json:"age"`
}


func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = bson.NewObjectId()

    err = session.DB("taskdb").C("users").Insert(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}


func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")

    if !bson.IsObjectIdHex(id) {
        http.Error(w, "invalid id format", http.StatusBadRequest)
        return
    }

    var user User
    err := session.DB("taskdb").C("users").FindId(bson.ObjectIdHex(id)).One(&user)
    if err != nil {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}


func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")

    if !bson.IsObjectIdHex(id) {
        http.Error(w, "invalid id format", http.StatusBadRequest)
        return
    }

    err := session.DB("taskdb").C("users").RemoveId(bson.ObjectIdHex(id))
    if err != nil {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}