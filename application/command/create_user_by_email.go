package command

import (
	"context"
	"errors"
	"github.com/palantir/stacktrace"
	domError "go-microservice/domain/error"
	"go-microservice/domain/service"
)

type CreateUserByEmailCommand struct {
	creatorService service.UserCreator
}

func NewCreateUserByEmailCommand(creatorService service.UserCreator) CreateUserByEmailCommand {
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
	serviceCredentials := service.UserCreationRequest{
		Firstname: credentials.Firstname,
		Surname:   credentials.Surname,
		Email:     credentials.Email,
		Password:  credentials.Password,
	}

	err := c.creatorService.Create(serviceCredentials)
	if err != nil {
		var emailErr domError.EmailAlreadyExists
		var passwordErr domError.PasswordTooShort

		switch {
		case errors.As(err, &emailErr):
			return stacktrace.Propagate(
				emailErr,
				"Failed to create user: Email '%s' is already in use.",
				emailErr.Email,
			)
		case errors.As(err, &passwordErr):
			return stacktrace.Propagate(
				passwordErr,
				"Failed to create user: Password is too short, it must be at least %d characters.",
				passwordErr.MinLength,
			)
		default:
			return stacktrace.Propagate(
				err,
				"Failed to execute CreateUserByEmailCommand.",
			)
		}
	}

	return nil
}
