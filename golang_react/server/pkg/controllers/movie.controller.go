package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/models"
	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/utils"
	"github.com/gorilla/mux"
)

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	id, err := strconv.Atoi(movieId)
	if err != nil {
		fmt.Println("error while parsing")
		utils.ErrorJSON(w, err)
		return
	}

	fmt.Println("Movie ID", id)

	movie := models.Movie{
		ID:          id,
		Title:       "The Da Vinci COde",
		Description: "SOme description",
		Year:        2016,
		ReleaseDate: time.Date(2016, 11, 9, 1, 0, 0, 0, time.Local),
		RunTime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = utils.WriteJSON(w, http.StatusOK, movie, "movie")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {}
