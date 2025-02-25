package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

// Internal server error
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	// URI: uniform resource identifier (both for abstact and physical resource)
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
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

func (app *application) render(w http.ResponseWriter, r *http.Request,
	status int, page string, data templateData) {

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	// Check for runtime rendering html errors, first write to a buffer, then if no errors
	// copy the buffer into the response.
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}


func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}