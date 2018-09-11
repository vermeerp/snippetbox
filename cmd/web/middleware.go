package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// LogRequest logs requests
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pattern := `%s - "%s %s %s"`
		log.Printf(pattern, r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})

}

// SecureHeaders sets headers for security features
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}

		next.ServeHTTP(w, r)
	})
}

// RequireLogin guards routes that should only be accessed after login
func (app *App) RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the app.LoggedIn() helper to get the status for the current user.
		loggedIn, err := app.LoggedIn(r)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		// If they are not logged in, redirect them to the login page and return
		// from the middleware chain so that no subsequent handlers in the chain
		// are executed.
		if !loggedIn {
			http.Redirect(w, r, "/user/login", 302)
			return
		}

		// Otherwise call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

// NoSurf middleware function which uses a customized CSRF cookie with
// the Secure, Path and HttpOnly flags set.
func NoSurf(next http.HandlerFunc) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
