package main

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/bzholmyrza/SPA-Final/internal/data"
)

func (app *application) createMusicHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMusicHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	music := data.Music{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Dark Side",
		Runtime:   102,
		Genres:    []string{"rock", "pop"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"music": music}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
