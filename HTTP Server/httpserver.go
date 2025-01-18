package main

import (
	"fmt"
	"net/http"
)

// Processing dynamic requests
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HTTP Server running...")
	fmt.Fprintf(w, "%s\n", r.URL.Query().Get("token"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")

	// Serving static assets
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Accept connections
	http.ListenAndServe(":8080", nil)
}
