package main

import (
	"fmt"
	"net/http"
)

func main()  {	
	// HTTP request multiplexer to match url of requests
	router := http.NewServeMux()

	// sample route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request at /")
		fmt.Fprintf(w,"Hello world!")
	})

	// set server and start
	server:= http.Server{
		Addr: ":8080",
		Handler: router,
	}
	fmt.Println("Starting server on :8080...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}