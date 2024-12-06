package todo

import (
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
			detail: "Your task goes here!",
		},
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorld received request at /")
	// write to the response which returns to client
	fmt.Fprintf(w, "Hello world!") 
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetItem received request")
	// print out the task details
	for i := 0; i < len(Tasks); i++ {
		fmt.Fprintf(w, "Task number %d: %s", i, Tasks[i].detail)
	}
}
