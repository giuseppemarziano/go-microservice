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

func SetupHTTPServer(cfg *Config) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServerAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
