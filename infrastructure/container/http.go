package container

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

func SetupHTTP() *http.Client {
	timeoutStr := os.Getenv("HTTP_CLIENT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil || timeout <= 0 {
		timeout = 30 // default timeout
	}
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	return client
}

func (c *Container) SetupHTTPServer(router http.Handler) *http.Server {
	addr := os.Getenv("HTTP_SERVER_ADDR")
	if addr == "" {
		addr = ":8080" // default address
	}
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
