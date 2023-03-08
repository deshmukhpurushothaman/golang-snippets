package routes

import (
	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/controllers"
	"github.com/gorilla/mux"
)

var Routes = func(r *mux.Router) {
	r.HandleFunc("/status", controllers.GetBook).Methods("GET")
}
