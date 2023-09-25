package main

import (
	"context"
	"go-microservice/infrastructure/container"
	"log"
)

func main() {
	ctx := context.Background()

	c, err := container.NewContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	server := c.HTTPServer
	if server == nil {
		log.Fatal("Server not initialized")
	}

	log.Printf("Starting server on %s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
