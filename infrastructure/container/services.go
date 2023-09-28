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
	GetUserAuthenticatorService(ctx context.Context) service.UserAuthenticator

	GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand
	GetCreateUserAuthenticationCommand(ctx context.Context) command.AuthenticateUserCommand

	GetGetAllUsersQuery(ctx context.Context) query.GetAllUsersQuery
	GetGetUserByEmailQuery(ctx context.Context) query.GetUserByEmailQuery
	GetGetUserByUUIDQuery(ctx context.Context) query.GetUserByUUIDQuery
}

// Services

func (c *Container) GetCreateUserService(ctx context.Context) service.UserCreator {
	return service.NewCreatorService(ctx, c.GetUserRepository(ctx), c.config.BCryptCost)
}

func (c *Container) GetUserAuthenticatorService(ctx context.Context) service.UserAuthenticator {
	return service.NewUserAuthenticationService(ctx, c.GetUserRepository(ctx))
}

// Repositories

func (c *Container) GetUserRepository(ctx context.Context) repositories.UserRepository {
	return mysql.NewUserRepository(ctx, c.db)
}

// Commands

func (c *Container) GetCreateUserByEmailCommand(ctx context.Context) command.CreateUserByEmailCommand {
	return command.NewCreateUserByEmailCommand(c.GetCreateUserService(ctx))
}

func (c *Container) GetCreateUserAuthenticationCommand(ctx context.Context) command.AuthenticateUserCommand {
	return command.NewAuthenticateUserCommand(c.GetUserAuthenticatorService(ctx))
}

// Queries

func (c *Container) GetGetAllUsersQuery(ctx context.Context) query.GetAllUsersQuery {
	return query.NewGetAllUsersQuery(c.GetUserRepository(ctx))
}

func (c *Container) GetGetUserByEmailQuery(ctx context.Context) query.GetUserByEmailQuery {
	return query.NewGetUserByEmailQuery(c.GetUserRepository(ctx))
}

func (c *Container) GetGetUserByUUIDQuery(ctx context.Context) query.GetUserByUUIDQuery {
	return query.NewGetUserByUUIDQuery(c.GetUserRepository(ctx))
}
