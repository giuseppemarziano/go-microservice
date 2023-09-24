package container

import (
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
)

type contextKey string

const ContainerKey contextKey = "Container"

type Container struct {
	HTTPClient *http.Client
	HTTPServer *http.Server
	DB         *gorm.DB
	Config     Config
	Services   Services
}

func NewContainer(ctx context.Context, router http.Handler) (*Container, error) {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConnection, err := NewGormDBConnection(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	httpClient := SetupHTTPClient(cfg)
	httpServer := SetupHTTPServer(cfg, router)

	return &Container{
		HTTPClient: httpClient,
		HTTPServer: httpServer,
		DB:         dbConnection,
		Config:     *cfg,
	}, nil
}
