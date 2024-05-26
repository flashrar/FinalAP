package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// movie routes here
	//router.HandlerFunc(http.MethodGet, "/v1/movies", app.requireActivatedUser(app.listMoviesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	// user routes here
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users", app.requireActivatedUser(app.getAllUserInfoHandler))
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", app.requireAdminUser(app.deleteUserInfoHandler))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.requireActivatedUser(app.getUserInfoHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/users/:id", app.requireAdminUser(app.editUserInfoHandler))

	// Return the httprouter instance.
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	// wrapping the router with rateLimiter() middleware to limit requests' frequency
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
