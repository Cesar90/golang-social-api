package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"status": "OK"}`))
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	// if err := writeJSON(w, http.StatusOK, post); err != nil {
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		// log.Print(err.Error())
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)
	}
}
