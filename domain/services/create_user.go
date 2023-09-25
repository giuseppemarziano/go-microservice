package services

import (
	"context"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
)

type Creator interface {
	Create(ctx context.Context, credentials UserRegistrationRequest) error
}

type creator struct {
	userRepository repositories.UserRepository
}

func NewCreatorService(ctx context.Context, repository repositories.UserRepository) Creator {
	return &creator{
		userRepository: repository,
	}
}

type UserRegistrationRequest struct {
	Username string
	Email    string
	Password string
}

func (c *creator) Create(ctx context.Context, credentials UserRegistrationRequest) error {
	credentialsData := entities.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	err := c.userRepository.CreateUser(ctx, &credentialsData)
	if err != nil {
		return stacktrace.Propagate(err, "error on creating user with email: %s", credentials.Email)
	}

	return nil
}
