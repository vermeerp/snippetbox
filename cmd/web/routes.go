package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
)

// Change the signature so we're returning a http.Handler instead of a
// *http.ServeMux.
func (app *App) Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/new", http.HandlerFunc(app.NewSnippet))
	mux.Post("/snippet/new", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet)) // Moved downwards

	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
