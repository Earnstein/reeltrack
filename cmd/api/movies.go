package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/earnstein/reeltrack/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create movie handler")
}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
		CreatedAt: time.Now(),
	}
	response := envelope{"movie": movie}

	err = app.writeJSON(w, http.StatusOK, response, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server could not process the request", http.StatusInternalServerError)
		return
	}
}
