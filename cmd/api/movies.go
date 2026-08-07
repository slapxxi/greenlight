package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/slapxxi/greenlight/internal/data"
)

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string   `json:"title"`
		Year   int32    `json:"year"`
		Rutime int32    `json:"runtime"`
		Genres []string `json:"genres"`
	}
	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
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
