package main

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Internal server error
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	// URI: uniform resource identifier (both for abstact and physical resource)
	var (
		method = r.Method
		uri = r.URL.RequestURI()
		// Stack trace of the current go-routine
		trace = string(debug.Stack())
	)
	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri),
		slog.String("trace", trace)) 
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}