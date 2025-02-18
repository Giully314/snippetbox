package main

import (
	"log"
	"net/http"
)


func main() {
	// Create a router for the application (usually there is only one)
	// which is responsible for dispatch the requests.
	// Method based routing allows to specify the type of http request to 
	// handle.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	log.Print("Starting server on :4000")

	// Use the default http server for this basic example.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
