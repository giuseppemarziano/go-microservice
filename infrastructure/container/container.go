package container

import (
	"context"
	"errors"
	"net/http"
	"os"
)

type contextKey string

const ContainerKey contextKey = "Container"

type Container struct {
	Database interface{}
	Http     *http.Client
	Config   Config
}

func NewContainer(ctx context.Context) (*Container, error) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		return nil, errors.New("database DSN is not set")
	}

	dbConnection := NewGormDBConnection(dsn)
	httpClient := SetupHTTP()
	config := NewConfig()

	return &Container{
		Database: dbConnection,
		Http:     httpClient,
		Config:   config,
	}, nil
}
