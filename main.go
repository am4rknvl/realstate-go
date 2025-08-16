package main

import (
	"log"
	"net/http"

	"github.com/ama4rknvl/realestate-go/database"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to DB
	database.Connect()

	// Create router
	r := mux.NewRouter()

	// Simple test endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running!"))
	}).Methods("GET")

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
