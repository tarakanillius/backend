// main.go
package main

import (
	"log"
	"my-app/internal/utils"
	"my-app/routes"
	"net/http"
)

func main() {
  // Initialize Firebase
  utils.InitFirebase()

	// Connect to the database
	if err := utils.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Setup routes
	r := routes.SetupRoutes()

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
