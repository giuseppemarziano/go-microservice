package container

import (
	"context"
	"go-microservice/application/command"
	_ "go-microservice/domain/entities"
	"go-microservice/domain/repositories"
	"go-microservice/domain/services"
	"go-microservice/infrastructure/persistence/mysql"
)

var _ Services = &Container{}

type Services interface {
	GetUserRepository(ctx context.Context) repositories.UserRepository
	GetRegisterUserService(ctx context.Context) services.Creator
	GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand
}

func (c *Container) GetRegisterUserService(ctx context.Context) services.Creator {
	return services.NewCreatorService(ctx, c.GetUserRepository(ctx))
}

func (c *Container) GetUserRepository(ctx context.Context) repositories.UserRepository {
	return mysql.NewUserRepository(c.DB)
}

func (c *Container) GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand {
	return command.NewCreateUserByEmailCommand(c.GetRegisterUserService(ctx))
}
