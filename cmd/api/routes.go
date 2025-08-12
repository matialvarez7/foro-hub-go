package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/topics", app.createTopicHandler)
	router.HandlerFunc(http.MethodGet, "/v1/topics/:id", app.showTopichandler)

	return router
}
