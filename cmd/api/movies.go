package main

import (
	"fmt"
	"net/http"
	"time"

	"codingmoon.io/go-test/internal/data"
)

// Post Movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Get Movie
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// read id param
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		http.NotFound(w, r)

		return
	}

	// Movie struct 사용
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// JSON response 작성
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)

	// server 에러
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
