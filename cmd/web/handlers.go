package main

import (
	"net/http"
	"strconv"
)

// Home function  writes a plain-text "Hello from Snippetbox"
// message as the HTTP response body.
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	// Fetch a slice of the latest snippets from the database.
	snippets, err := app.Database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Pass the slice of snippets to the "home.page.html" templates.
	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})
}

// ShowSnippet handler function.
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if snippet == nil {
		app.NotFound(w)
		return
	}

	// Render the show.page.html template, passing in the snippet data wrapped in
	// our HTMLData struct.
	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet,
	})
}

// NewSnippet handler function.
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
