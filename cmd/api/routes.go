package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/pinatas", app.listPinatasHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pinatas", app.createPinataHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pinatas/:id", app.showPinataHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/pinatas/:id", app.updatePinataHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/pinatas/:id", app.deletePinataHandler)

	return app.recoverPanic(router)
}
