package container

import (
	"net/http"
	"time"
)

// SetupHTTPClient initializes and returns a new HTTP client with a configured timeout.
func SetupHTTPClient(cfg *Config) *http.Client {
	return &http.Client{
		Timeout: time.Second * time.Duration(cfg.HTTPClientTimeout), // set timeout from configuration
	}
}

// SetupHTTPServer initializes and returns a new HTTP server with configured address and timeouts.
func SetupHTTPServer(cfg *Config) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServerAddr, // set server address from configuration
		ReadTimeout:  5 * time.Second,    // set read timeout
		WriteTimeout: 10 * time.Second,   // set write timeout
		IdleTimeout:  15 * time.Second,   // set idle timeout
	}
}
