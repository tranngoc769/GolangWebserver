package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}
func main() {
	r := mux.NewRouter()
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/{title}/page/{page}", AllBooks)
	fmt.Print("Running 8080\n")
	http.ListenAndServe(":8080", r)
}
