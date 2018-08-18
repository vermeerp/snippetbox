package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new serve mux. Then use
	// the mux.HandleFunc() method to register our Home function as the handler for
	// the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// As before, the path given to the http.Dir function is relative to our project
	// repository root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the
	// handler for all URL paths that start with "/static/". For matching
	// paths, we strip the "/static" prefix before the request reaches the file
	// server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the serve mux we just created. If ListenAndServe() returns an error we
	// use the log.Fatal() function to record the error message and exit the program.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
