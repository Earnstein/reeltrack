package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	stat := envelope{
		"status": "available",
		"system-info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, stat, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server could not process the request", http.StatusInternalServerError)
		return
	}
}
