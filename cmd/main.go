package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akaroth/go-rest-api/config"
	"github.com/akaroth/go-rest-api/pkg/handlers"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/api/health", handlers.HeathCheckHandler).Methods("GET")
	fmt.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
