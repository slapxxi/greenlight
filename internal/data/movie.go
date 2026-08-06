package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Genres    []string  `json:"genres,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Title     string    `json:"title"`
	Version   int32     `json:"version"`
	Year      int32     `json:"year,omitempty"`
}
