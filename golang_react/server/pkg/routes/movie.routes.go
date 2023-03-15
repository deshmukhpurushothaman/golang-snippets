package routes

import (
	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/controllers"
	// "github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/cmd/api/main"
)

func (app *application) MovieRoutes() {
	app.router.HandleFunc("/v1/movie/{movieId}", controllers.GetOneMovie).Methods("GET")
	app.router.HandleFunc("/v1/movies", controllers.GetAllMovies).Methods("GET")
}
