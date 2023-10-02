package main

import (
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/route"
	"log"
)

func main() {
	// initialize dependency injection container
	c, err := container.NewContainer()
	if err != nil {
		log.Fatalf("failed to create container: %v", err)
	}

	// retrieve http server from container and set up routes
	server := c.HTTPServer
	if server == nil {
		log.Fatal("server not initialized")
	}
	server.Handler = route.Routes(*c)

	// start the http server
	log.Printf("starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
