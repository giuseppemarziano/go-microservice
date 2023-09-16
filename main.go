package main

import (
	"context"
	"log"
	"net/http"

	"go-microservice/infrastructure/container"
)

func main() {
	// Initialize the container
	ctx := context.Background()
	c := container.NewContainer(ctx)

	// Create a simple HTTP router (this should be more complex in a real app)
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Get the configured HTTP server from the container
	server := c.SetupHTTPServer(ctx, router)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
