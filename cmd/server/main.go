package main

import (
	"fmt"
	"log"
	"net/http"

	"gowire/internal/handlers"
	"gowire/internal/router"
)

func main() {
	h, err := handlers.NewHandlers()
	if err != nil {
		log.Fatalf("Failed to create handlers: %v", err)
	}

	mux := router.SetupRoutes(h)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
