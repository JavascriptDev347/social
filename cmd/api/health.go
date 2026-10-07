package main

import (
	"net/http"
)

var version string = "0.0.1"

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "OK",
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
