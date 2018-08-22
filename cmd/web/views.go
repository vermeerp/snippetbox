package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/vermeerp/snippetbox/pkg/models" // New import
)

// HTMLData struct acts as a wrapper for the dynamic data we want
// to pass to our templates. For now this just contains the snippet data that we
// want to display, which has the underling type *models.Snippet.
type HTMLData struct {
	Form     interface{}
	Path     string
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

// Create a humanDate function which returns a nicely formated string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// RenderHTML renders the HTML
func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {
	// If no data has been passed in, initialize a new empty HTMLData object.
	if data == nil {
		data = &HTMLData{}
	}

	// Add the current request URL path to the data.
	data.Path = r.URL.Path

	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	// Initialize a template.FuncMap object. This is essentially a string-keyed map
	// which acts as a lookup between the names of our custom template functions and
	// the functions themselves.
	fm := template.FuncMap{
		"humanDate": humanDate,
	}

	ts, err := template.New("").Funcs(fm).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err) // Use the new app.ServerError() helper.
		return
	}
	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter. If there's an error, call our error handler and then
	// return.
	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Write the contents of the buffer to the http.ResponseWriter. Again, this
	// is another time where we pass our http.ResponseWriter to a function that
	// takes an io.Writer.
	buf.WriteTo(w)
}
