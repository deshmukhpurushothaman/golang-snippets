package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type CurrentStatus struct {
	Status      string
	Environment string
	Version     string
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	currentStatus := CurrentStatus{
		Status:      "Available",
		Environment: "dev",
		Version:     "1.0.0",
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
