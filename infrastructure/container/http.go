package container

import (
	"net/http"
	"time"
)

func SetupHTTPClient(cfg *Config) *http.Client {
	return &http.Client{
		Timeout: time.Second * time.Duration(cfg.HTTPClientTimeout),
	}
}

func SetupHTTPServer(cfg *Config, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServerAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
