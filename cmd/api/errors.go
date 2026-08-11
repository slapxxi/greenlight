package main

import (
	"fmt"
	"net/http"
)

func (a *application) logError(r *http.Request, err error) {
	a.logger.Print(err)
}

func (a *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, msg any) {
	env := envelope{"error": msg}
	err := a.writeJSON(w, status, env, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.logError(r, err)
	message := "The sever encountered an error"
	a.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource was not found"
	a.errorResponse(w, r, http.StatusNotFound, message)
}

func (a *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The method %s is not allowed", r.Method)
	a.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (a *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	a.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
