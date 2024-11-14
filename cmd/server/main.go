package main

import (
	"fmt"
	"log"
	"net/http"

	"gowire/internal/router"
)

func main() {
	mux, err := router.SetupRoutes()
	if err != nil {
		log.Fatalf("Failed to setup routes: %v", err)
	}

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
