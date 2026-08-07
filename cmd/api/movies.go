package main

import (
	"fmt"
	"net/http"
	"time"

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

	fmt.Fprintf(w, "%+v\n", input)
}

func (a *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIdParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Genres:    []string{"drama", "comedy"},
		Runtime:   120,
		Title:     "The Matrix",
		Version:   1,
		Year:      1999,
	}
	err = a.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
