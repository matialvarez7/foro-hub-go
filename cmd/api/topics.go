package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/matialvarez7/foro-hub-go/internal/data"
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

	topic := data.Topic{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Topico de prueba",
		Message:   "Este es un tópico de prueba",
		Status:    "Abierto",
		AuthorID:  1,
		CourseID:  1,
		Replies: []data.Reply{
			{
				ID:         1,
				CreatedAt:  time.Now(),
				Message:    "Una respuesta para el tópico",
				TopicID:    id,
				AuthorID:   2,
				IsSolution: false,
			},
		},
	}

	err = app.writeJSON(w, http.StatusOK, topic, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
