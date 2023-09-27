package main

import (
	"context"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/routes"
	"log"
)

func main() {
	// Initialize container
	c, err := container.NewContainer(context.Background())
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	// Get the HTTP server from the container and set the Echo instance as its Handler
	server := c.HTTPServer
	if server == nil {
		log.Fatal("Server not initialized")
	}
	server.Handler = routes.SetupRoutes(*c)

	// Start the server
	log.Printf("Starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
