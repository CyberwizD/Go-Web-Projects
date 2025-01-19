package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {
	// Creating a new router
	r := mux.NewRouter()

	// Registering a request handler
	r.HandleFunc("/books/{title}/page/{page}", handler)

	// Setting the HTTP server's router
	http.ListenAndServe(":8080", r)
}

/* Features of gorilla/mux Router:
Methods - Restrict the request handler to specific HTTP methods.
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")

Hostnames & Subdomains - Restrict the request handler to specific hostnames or subdomains.
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

Schemes - Restrict the request handler to http/https.
r.HandleFunc("/secure", SecureHandler).Schemes("https")

Path Prefixes & Subrouters - Restrict the request handler to specific path prefixes.
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
*/
