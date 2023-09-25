package command

import (
	"context"
	"go-microservice/domain/entities"
	"go-microservice/domain/services"
)

type CreateUserByEmailCommand struct {
	creatorService services.Creator
}

func NewCreateUserByEmailCommand(creatorService services.Creator) CreateUserByEmailCommand {
	return CreateUserByEmailCommand{
		creatorService: creatorService,
	}
}

func (c *CreateUserByEmailCommand) Do(ctx context.Context, credentials entities.User) error {
	err := c.creatorService.Create(ctx, credentials)
	if err != nil {
		return err
	}

	return nil
}
