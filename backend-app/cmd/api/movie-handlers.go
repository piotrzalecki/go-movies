package main

import (
	"errors"
	"movies-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJson(w, err)
		return
	}

	movie := models.Movie{
		ID:          id,
		Title:       "Poranek kojota",
		Description: "Super fainy film",
		Year:        2000,
		ReleaseDate: time.Date(2000, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     110,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_ = app.writeJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

}
