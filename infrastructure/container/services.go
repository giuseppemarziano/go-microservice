package container

import (
	"context"
	_ "go-microservice/domain/entities"
	"go-microservice/domain/repositories"
	"go-microservice/domain/services"
	"go-microservice/infrastructure/persistence/mysql"
)

var _ Services = &Container{}

type Services interface {
	GetUserRepository(ctx context.Context) repositories.UserRepository
	GetCreateUserService(ctx context.Context) services.Creator
}

func (c *Container) GetCreateUserService(ctx context.Context) services.Creator {
	return services.NewCreatorService(ctx, c.GetUserRepository(ctx))
}

func (c *Container) GetUserRepository(ctx context.Context) repositories.UserRepository {
	return mysql.NewUserRepository(c.DB)
}
