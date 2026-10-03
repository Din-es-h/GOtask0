package main

import "fmt"

// Concept 1: Struct
type Task struct {
	ID        int
	Title     string
	Completed bool
}

type ScheduleTracker struct {
	tasks []Task // Concept 2: Slice
}

// Concept 3: Function with Error Return
func (st *ScheduleTracker) AddTask(id int, title string) error {
	if title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	st.tasks = append(st.tasks, Task{ID: id, Title: title, Completed: false})
	return nil
}

func (st *ScheduleTracker) MarkDone(id int) error {
	// Concept 4: Loop
	for i := range st.tasks {
		if st.tasks[i].ID == id {
			st.tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Concept 5: Goroutine (uses a unbuffered channel for sync instead of sync package)
func (st *ScheduleTracker) CalculateStatsAsync(ch chan map[string]int) {
	// Concept 6: Map
	stats := map[string]int{
		"total":     len(st.tasks),
		"completed": 0,
		"pending":   0,
	}

	for _, task := range st.tasks {
		if task.Completed {
			stats["completed"]++
		} else {
			stats["pending"]++
		}
	}

	ch <- stats // Send result back through channel
}

func main() {
	tracker := &ScheduleTracker{}

	// Adding tasks (testing error handling)
	if err := tracker.AddTask(1, "Review Go concurrency patterns"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := tracker.AddTask(2, "Write CLI Schedule Tracker"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := tracker.AddTask(3, "Refactor code architecture"); err != nil {
		fmt.Println("Error:", err)
	}

	// Mark tasks complete
	_ = tracker.MarkDone(1)
	_ = tracker.MarkDone(2)

	// Channel handles synchronization directly—no sync/time needed
	statsChan := make(chan map[string]int)

	go tracker.CalculateStatsAsync(statsChan) // Goroutine launch

	// Blocking receive from channel synchronizes execution naturally
	stats := <-statsChan

	// Display Output
	fmt.Println("--- Schedule Tracker Summary ---")
	fmt.Printf("Total Tasks:     %d\n", stats["total"])
	fmt.Printf("Tasks Completed: %d\n", stats["completed"])
	fmt.Printf("Tasks Pending:   %d\n", stats["pending"])
}
