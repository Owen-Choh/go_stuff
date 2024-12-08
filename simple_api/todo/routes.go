package todo

import "net/http"

func SetUpHttpMux() *http.ServeMux{
	// HTTP request multiplexer to match url of requests
	router := http.NewServeMux()

	// sample route
	router.HandleFunc("/", HelloWorld)
	router.HandleFunc("/task/all", GetAllTasks)
	router.HandleFunc("/task/{id}", GetTaskByIndex)

	return router
}