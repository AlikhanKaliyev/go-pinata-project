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
	router.HandlerFunc(http.MethodGet, "/v1/pinatas", app.requirePermission("pinatas:read", app.listPinatasHandler))
	router.HandlerFunc(http.MethodPost, "/v1/pinatas", app.requirePermission("pinatas:read", app.createPinataHandler))
	router.HandlerFunc(http.MethodGet, "/v1/pinatas/:id", app.requirePermission("pinatas:write", app.showPinataHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/pinatas/:id", app.requirePermission("pinatas:write", app.updatePinataHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/pinatas/:id", app.requirePermission("pinatas:write", app.deletePinataHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
