package container

import (
	"context"
	"net/http"
	"time"
)

func (c *Container) SetupHTTP(ctx context.Context) *http.Client {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	return client
}

func (c *Container) SetupHTTPServer(ctx context.Context, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
