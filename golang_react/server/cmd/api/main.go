package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/routes"
	"github.com/gorilla/mux"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5000, "Server to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment")
	flag.Parse()

	fmt.Println("Running")

	// http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
	// 	currentStatus := AppStatus{
	// 		Status:      "Available",
	// 		Environment: cfg.env,
	// 		Version:     version,
	// 	}

	// 	js, err := json.MarshalIndent(currentStatus, "", "\t")

	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(js)
	// })

	r := mux.NewRouter()
	routes.Routes(r)
	routes.MovieRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), r)

	if err != nil {
		log.Println(err)
	}
}
