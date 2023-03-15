package routes

import (
	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/controllers"
	"github.com/gorilla/mux"

	// "github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/cmd/api/main"

	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/models"
)

type application struct {
	config struct{}
	router *mux.Router
	models models.Models
}

func (app *application) MovieRoutes() {
	app.router.HandleFunc("/v1/movie/{movieId}", controllers.GetOneMovie).Methods("GET")
	app.router.HandleFunc("/v1/movies", controllers.GetAllMovies).Methods("GET")
}
