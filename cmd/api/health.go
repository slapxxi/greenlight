package main

import (
	"net/http"
)

func (a *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":      "available",
		"environment": a.config.env,
		"version":     version,
	}
	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
