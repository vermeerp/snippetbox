package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home function  writes a plain-text "Hello from Snippetbox"
// message as the HTTP response body.
func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.page.html",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set (notice that we can pass the slice of file paths
	// as a variadic parameter). If there's an error, we log the detailed error
	// message and use the http.Error() function to send a generic 500 Internal
	// Server Error response.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Our template set contains three named templates: base, page-title and
	// page-body (note that every template in your template set must have a
	// unique name). We use the ExecuteTemplate() method to execute the "base"
	// template and write its content to our http.RespsonseWriter. The last
	// parameter to ExecuteTemplate() represents any dynamic data that we want to
	// pass in, which for now we'll leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// ShowSnippet handler function.
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it couldn't
	// be converted to an integer, or the value is less than 1, we return a 404
	// Not Found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
}

// NewSnippet handler function.
func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
