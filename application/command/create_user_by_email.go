package command

import (
	"context"
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

type UserRegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *CreateUserByEmailCommand) Do(ctx context.Context, credentials UserRegistrationRequest) error {
	err := c.creatorService.Create(ctx, credentials)
	if err != nil {
		return err
	}

	return nil
}
