package main

import (
	"context"
	"errors"
	"fmt"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/routes"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	router := routes.SetupRoutes()

	c, err := container.NewContainer(ctx, router)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	server := c.HTTPServer
	if server == nil {
		log.Fatal("Server not initialized")
	}

	addr := server.Addr
	fmt.Printf("Server Setup Complete: %+v\n", server)
	log.Printf("Starting server on %s\n", addr)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}
}
