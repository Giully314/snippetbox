// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	// Add a field in the header of the http request.
// 	w.Header().Add("Server", "Go")
// 	w.Write([]byte("Hello from snippetbox alpha."))
// }

// func snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	// msg := fmt.Sprintf("This is a snippet view for the user with ID %d", id)
// 	// w.Write([]byte(msg))
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprintf(w, `{"id": "%d"}`, id)
// }

// func snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Display a form for creating a custom snippet."))
// }

// func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusCreated)

// 	w.Write([]byte("Save a new snippet."))
// }



// func main() {
// 	// Create a router for the application (usually there is only one)
// 	// which is responsible for dispatch the requests.
// 	// Method based routing allows to specify the type of http request to 
// 	// handle.
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /{$}", home)
// 	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
// 	mux.HandleFunc("GET /snippet/create", snippetCreate)
// 	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

// 	// When a path ends with '/', it will match any path that doesn't exist
// 	// but has the same subtree root. For example the path /snippet/clown
// 	// will be matched by this (because /snippet/clown doesn't exist).
// 	// To prevent this, we can end the path with {$}. This will result 
// 	// in 404 error.
// 	// mux.HandleFunc("/snippet/{$}", snippet)

// 	// For overlapping methods, the most specific (like in C++ template)
// 	// take the precedence. 

// 	log.Print("Starting server on :4000")

// 	// Access header can be done using Header() (Header is just a map 
// 	// from string to []string). Go will automatically call the 
// 	// CanonicalMIMEHeaderKey in textproto to make it in the right format.

// 	// Use the default http server for this basic example.
// 	err := http.ListenAndServe(":4000", mux)
// 	log.Fatal(err)
// }
