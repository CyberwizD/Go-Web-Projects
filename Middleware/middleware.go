package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func goo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Go")
}

func Lang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Lang")
}

func main() {
	http.HandleFunc("/Go", logging(goo))
	http.HandleFunc("/Lang", logging(Lang))

	http.ListenAndServe(":8080", nil)
}
