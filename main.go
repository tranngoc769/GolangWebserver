package main

import (
	"fmt"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Print("Running :" + port + "\n")
	http.ListenAndServe(":"+port, r)
}
