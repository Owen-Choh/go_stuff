package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type task struct {
	detail string
}

var Tasks []task

func Init() {
	fmt.Println("Initialise todo")
	Tasks = []task{
		{
			detail: "Task number 1 goes here!",
		},
		{
			detail: "Your task goes here!",
		},
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorld received request at /")
	// write to the response which returns to client
	fmt.Fprintf(w, "Hello world!") 
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for all tasks")
	// print out the task details
	w.Write(getalltasks())
}

func getalltasks() []byte {
	var output string
	for i := 0; i < len(Tasks); i++ {
		output += fmt.Sprintf("Task number %d: %s", i, Tasks[i].detail)
	}
	return []byte(output)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write(gettask(id))
}

func gettask(id string) []byte {
	return []byte("received request for task: " + id)
}
