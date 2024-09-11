package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("GET /snippet/view/{id}", SnippetView)
	mux.HandleFunc("GET /snippet/create", SnippetCreate)
	mux.HandleFunc("POST /snippet/create", SnippetCreatePost)

	log.Print("Starting server at :4000")
	error := http.ListenAndServe(":4000", mux)

	if error != nil {
		log.Fatal(error)
	}
}
