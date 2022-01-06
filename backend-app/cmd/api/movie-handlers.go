package main

import (
	"errors"
	"net/http"
	"strconv"

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

	movie, err := app.models.DB.Get(id)

	if err != nil {
		app.logger.Print(err)
	}

	err = app.writeJson(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJson(w, err)
		return
	}
	err = app.writeJson(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMovieByGenre(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}
	movies, err := app.models.DB.All(genreID)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) editMovie(w http.ResponseWriter, r *http.Request) {
	type jsonResponse struct {
		OK bool `json:"ok"`
	}

	ok := jsonResponse{
		OK: true,
	}

	err := app.writeJson(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
