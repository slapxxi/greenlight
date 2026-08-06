package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/slapxxi/greenlight/internal/data"
)

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create movie")
}

func (a *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
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
		a.logger.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
