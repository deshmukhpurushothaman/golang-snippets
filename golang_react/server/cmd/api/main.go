package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/deshmukhpurushothaman/golang-snippets/golang_react/server/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://deshmukh:deshmukh@localhost/go_movies?sslmode=disable", "Postgres connection string")
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

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	routes.Routes(r)
	routes.MovieRoutes(r)
	http.Handle("/", r)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), r)

	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}
	return db, nil
}
