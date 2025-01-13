package main

import (
	"fmt"
	"net/http"

	"github.com/Owen-Choh/go_stuff/simple_api/todo"
)

func main()  {
	todo.Init()

	router := todo.SetUpHttpMux()

	router.HandleFunc("OPTIONS /cors", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("got cors options request")
		fmt.Println(r.Header)
		
		w.Header().Add("Access-Control-Allow-Origin","*")
		w.Header().Add("Access-Control-Allow-Methods","GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers","Content-Type")
		w.Header().Add("Access-Control-Max-Age","10000")
		fmt.Fprint(w, "test cors options request success")
	})

	router.HandleFunc("GET /cors", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("got cors GET request")
		fmt.Println(r.Header)
		w.Header().Add("Access-Control-Allow-Origin","*")
		w.Header().Add("Access-Control-Allow-Methods","GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers","Content-Type")
		w.Header().Add("Access-Control-Max-Age","10000")
		fmt.Fprint(w, "test cors get request success")
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