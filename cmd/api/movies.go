package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/slapxxi/greenlight/internal/data"
	"github.com/slapxxi/greenlight/internal/validator"
)

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: data.Runtime(input.Runtime),
		Genres:  input.Genres,
	}
	data.ValidateMovie(v, movie)
	if !v.Valid() {
		a.fieldValidationResponse(w, r, v.Errors)
		return
	}

	err = a.models.Movie.Insert(movie)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIdParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}
	movie, err := a.models.Movie.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
			return
		default:
			a.serverErrorResponse(w, r, err)
			return
		}
	}

	headers := make(http.Header)
	headers.Set("ETag", fmt.Sprintf("%d", movie.Version))

	err = a.writeJSON(w, http.StatusOK, envelope{"movie": movie}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
