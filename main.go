package main

import (
    "log"
    "net/http"
)

// Define a Home function which writes a plain-text "Hello from Snippetbox"
// message as the HTTP response body.
func Home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func main() {
    // Use the http.NewServeMux() function to initialize a new serve mux. Then use
    // the mux.HandleFunc() method to register our Home function as the handler for
    // the "/" URL pattern.
    mux := http.NewServeMux()
    mux.HandleFunc("/", Home)

    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":4000")
    // and the serve mux we just created. If ListenAndServe() returns an error we
    // use the log.Fatal() function to record the error message and exit the program.
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
