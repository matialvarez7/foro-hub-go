package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTopicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new topic")
}

func (app *application) showTopichandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details for the topic: %d\n", id)
}
