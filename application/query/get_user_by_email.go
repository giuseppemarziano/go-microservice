package query

import (
	"context"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
)

type GetUserByEmailQuery struct {
	userRepository repositories.UserRepository
}

func NewGetUserByEmailQuery(userRepository repositories.UserRepository) GetUserByEmailQuery {
	return GetUserByEmailQuery{
		userRepository: userRepository,
	}
}

func (gue *GetUserByEmailQuery) Do(ctx context.Context, email string) (*entities.User, error) {
	return gue.userRepository.GetUserByEmail(ctx, email)
}
