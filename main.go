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
	c, err := container.NewContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	e := routes.SetupRoutes()

	server := c.SetupHTTPServer(e)
	fmt.Printf("Server Setup Complete: %+v\n", server)

	addr := server.Addr
	log.Printf("Starting server on %s\n", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}
}
