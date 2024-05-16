package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akaroth/go-rest-api/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/health", handlers.HeathCheckHandler).Methods("GET")
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
