package container

import (
	"context"
)

type contextKey string

const ContainerKey contextKey = "Container"

type Container struct {
	DatabaseConnection interface{}
	HttpClient         interface{}
}

func NewContainer(ctx context.Context) *Container {
	//dbConnection := initDatabaseConnection()
	//httpClient := initHttpClient()

	return &Container{
		//DatabaseConnection: dbConnection,
		//HttpClient:         httpClient,
	}
}
