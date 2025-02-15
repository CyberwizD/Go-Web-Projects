package main

import "net/http"

// Serving static files; CSS, Javascript from a specific directory

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}
