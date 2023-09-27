package command

import (
	"context"
	"go-microservice/domain/service"
)

type CreateUserByEmailCommand struct {
	creatorService service.Creator
}

func NewCreateUserByEmailCommand(creatorService service.Creator) CreateUserByEmailCommand {
	return CreateUserByEmailCommand{
		creatorService: creatorService,
	}
}

type UserRegistrationRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (c *CreateUserByEmailCommand) Do(ctx context.Context, credentials UserRegistrationRequest) error {
	err := c.creatorService.Create(service.UserRegistrationRequest(credentials))
	if err != nil {
		return err
	}

	return nil
}
