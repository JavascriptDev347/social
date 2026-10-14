package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Checks the API health
//	@Description	Returns the API status and version
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	error
//	@Router			/health [get]
func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "OK",
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
