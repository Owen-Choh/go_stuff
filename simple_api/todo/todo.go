package todo

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request at /")
	fmt.Fprintf(w,"Hello world!") // write to the response which returns to client
}
