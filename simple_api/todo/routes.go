package todo

import "net/http"

func SetUpHttpMux() *http.ServeMux{
	// HTTP request multiplexer to match url of requests
	router := http.NewServeMux()

	// sample route
	router.HandleFunc("GET /", HelloWorld)
	router.HandleFunc("GET /task/all", GetAllTasks)
	router.HandleFunc("GET /task/{id}", GetTaskByIndex)

	return router
}