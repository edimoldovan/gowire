package main

import (
	"fmt"
	"log"
	"net/http"

	"gowire/internal/handlers"
	"gowire/internal/router"
)

func main() {
	// Initialize handlers
	h := handlers.NewHandlers()

	// Setup routes
	mux := router.SetupRoutes(h)

	// Define the port
	port := ":8080"

	// Start the server
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
