package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Detail string
}

var Tasks []Task

func Init() {
	fmt.Println("Initialise todo")
	Tasks = []Task{
		{
			Detail: "Task number 1 goes here!",
		},
		{
			Detail: "Your task goes here!",
		},
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(Tasks)
	if err != nil {
		log.Fatalf("Error marshalling data to json: %v", Tasks)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
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
		output += fmt.Sprintf("Task number %d: %s", i, Tasks[i].Detail)
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
