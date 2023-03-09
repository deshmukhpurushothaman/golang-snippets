package routes

import (
	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/controllers"
	"github.com/gorilla/mux"
)

var MovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/movie/{movieId}", controllers.GetOneMovie).Methods("GET")
	router.HandleFunc("/v1/movies", controllers.GetAllMovies).Methods("GET")
}
