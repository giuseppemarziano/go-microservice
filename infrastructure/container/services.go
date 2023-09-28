package container

import (
	"context"
	"go-microservice/application/command"
	"go-microservice/application/query"
	_ "go-microservice/domain/entities"
	"go-microservice/domain/repositories"
	"go-microservice/domain/service"
	"go-microservice/infrastructure/persistence/mysql"
)

var _ Services = &Container{}

type Services interface {
	GetUserRepository(ctx context.Context) repositories.UserRepository

	GetCreateUserService(ctx context.Context) service.UserCreator

	GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand
}

// Services

func (c *Container) GetCreateUserService(ctx context.Context) service.UserCreator {
	return service.NewCreatorService(ctx, c.GetUserRepository(ctx))
}

// Repositories

func (c *Container) GetUserRepository(ctx context.Context) repositories.UserRepository {
	return mysql.NewUserRepository(ctx, c.DB)
}

// Commands

func (c *Container) GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand {
	return command.NewCreateUserByEmailCommand(c.GetCreateUserService(ctx))
}

// Queries

func (c *Container) GetGetAllUsersQuery(ctx context.Context) query.GetAllUsersQuery {
	return query.NewGetAllUsersQuery(c.GetUserRepository(ctx))
}
