package container

import (
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
)

type Container struct {
	HTTPClient *http.Client
	HTTPServer *http.Server
	DB         *gorm.DB
	Config     Config
}

type RouteHandler interface {
	SetupRoutes() http.Handler
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

	return &Container{
		HTTPClient: SetupHTTPClient(cfg),
		HTTPServer: SetupHTTPServer(cfg),
		DB:         dbConnection,
		Config:     *cfg,
	}, nil
}
