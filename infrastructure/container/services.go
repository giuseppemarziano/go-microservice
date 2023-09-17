package container

import (
	"context"
	"go-microservice/domain/service"
)

var _ Services = &Container{}

type Services interface {
	GetHelloWorldService(ctx context.Context) service.TestService
}

func (c *Container) GetHelloWorldService(ctx context.Context) service.TestService {
	return service.NewTestService()
}
