package main

import (
	"fmt"
	"net/http"
)

/* Registering a request handler
The function receives two parameters:

An http.ResponseWriter which is where you write your text/html response to.
An http.Request which contains all information about this HTTP request including things like the URL or header fields. */

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, There! \n")
	fmt.Fprintf(w, "you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
