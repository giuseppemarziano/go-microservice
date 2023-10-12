package container

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
)

// Container holds application-wide instances like HTTP client, server, database connection, and configuration.
type Container struct {
	HTTPClient *http.Client
	HTTPServer *http.Server
	db         *gorm.DB
	config     Config
}

// RouteHandler interface defines a method for setting up HTTP routes.
type RouteHandler interface {
	SetupRoutes() http.Handler
}

// NewContainer initializes a new Container with all dependencies set up.
func NewContainer() (*Container, error) {
	// load application configuration
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// establish database connection using individual DSN components from the loaded configuration
	dbConnection, err := NewGormDBConnection(
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseParams,
	)
	if err != nil {
		return nil, err
	}

	// initialize and return a new Container with all dependencies
	return &Container{
		HTTPClient: SetupHTTPClient(cfg),
		HTTPServer: SetupHTTPServer(cfg),
		db:         dbConnection,
		config:     *cfg,
	}, nil
}
