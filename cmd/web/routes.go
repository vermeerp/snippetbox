package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
)

// Routes handles routing the request
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
