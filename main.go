package main

import (
	"context"
	"fmt"
	_ "github.com/labstack/echo/v4" // import echo
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/routes" // import your routes
	"log"
	_ "net/http"
)

func main() {
	// Initialize the container
	ctx := context.Background()
	c := container.NewContainer(ctx)

	// Setup Echo routes (replace this with your actual routes setup function)
	e := routes.SetupRoutes()

	// Get the configured HTTP server from the container
	server := c.SetupHTTPServer(ctx, e) // Pass Echo router here
	fmt.Println(server)
	// Start the HTTP server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080")) // Use Echo's start method
}
