package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s path: %s. what actually happen here: %s", r.Method, r.URL.Path, err.Error())

	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error: %s path: %s. what actually happen here: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}
func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflic request error: %s path: %s. what actually happen here: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorf("conflict", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error: %s path: %s. what actually happen here: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("Not found error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusNotFound, "Resource not found")
}
