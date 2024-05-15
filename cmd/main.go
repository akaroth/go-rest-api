package main

import (
	"github.com/akaroth/go-rest-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
}
