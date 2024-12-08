package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func jsonResponse(w http.ResponseWriter, payload any) {
	data, err := json.Marshal(payload)
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
	// returns the all tasks details
	fmt.Println("received request for all tasks")

	jsonResponse(w, Tasks)
}

func GetTaskByIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for specific task")

	id := r.PathValue("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if index >= len(Tasks) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse(w, Tasks[index])
}
