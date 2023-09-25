package container

import (
	"context"
	"github.com/labstack/gommon/log"
	"go-microservice/infrastructure/routes"
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
}

func NewContainer(ctx context.Context) (*Container, error) {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConnection, err := NewGormDBConnection(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	httpClient := SetupHTTPClient(cfg)
	httpServer := SetupHTTPServer(cfg, routes.SetupRoutes())

	return &Container{
		HTTPClient: httpClient,
		HTTPServer: httpServer,
		DB:         dbConnection,
		Config:     *cfg,
	}, nil
}
