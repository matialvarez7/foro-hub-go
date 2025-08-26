package main

import (
	"fmt"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	err := app.writeJSON(w, status, message, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	e := Error{Error: message}
	app.errorResponse(w, r, http.StatusInternalServerError, e)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	e := Error{Error: message}
	app.errorResponse(w, r, http.StatusNotFound, e)
}

func (app *application) mehtodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s mehtod is not supported for this resource", r.Method)
	e := Error{Error: message}
	app.errorResponse(w, r, http.StatusMethodNotAllowed, e)
}
