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

// hello world!
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorld received request at /")
	// write to the response which returns to client
	fmt.Fprintf(w, "Hello world!") 
}

// returns the all tasks details
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for all tasks")

	jsonResponse(w, Tasks)
}

// returns the task at the index
func GetTaskByIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for specific task")

	id := r.PathValue("taskId")
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

// deletes task while keeping original order
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request to delete task")

	requestedIndex := r.PathValue("taskId")
	index, err := strconv.Atoi(requestedIndex)

	if err != nil || index < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if index >= len(Tasks) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Tasks = append(Tasks[:index], Tasks[index + 1:]...)
}

// insert the task at the index specified, at end of list if index is too large
func CreateTaskAtIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request to add task at index")

	// get the index from request
	id := r.PathValue("taskId")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 {
		fmt.Println("invalid index received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if index > len(Tasks) {
		index = len(Tasks)
	}

	// get the task to put into the list
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var inputTask Task

	err = decoder.Decode(&inputTask)
	if err != nil {
		fmt.Println("invalid task received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if inputTask.Detail == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	Tasks = append(Tasks, Task{})
	copy(Tasks[index+1:], Tasks[index:])
	Tasks[index] = inputTask

	w.WriteHeader(http.StatusCreated)
}
