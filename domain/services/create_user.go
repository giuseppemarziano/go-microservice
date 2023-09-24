package services

import (
	"context"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
)

type Creator interface {
	Create(ctx context.Context, credentials entities.User) error
}

type creator struct {
	userRepository repositories.UserRepository
}

func NewCreatorService(ctx context.Context, repository repositories.UserRepository) Creator {
	return &creator{
		userRepository: repository,
	}
}

func (c *creator) Create(ctx context.Context, credentials entities.User) error {
	err := c.userRepository.CreateUser(ctx, &credentials)
	if err != nil {
		return stacktrace.Propagate(err, "error on creating user with email: %s", credentials.Email)
	}

	return nil
}
