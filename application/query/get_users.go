package query

import (
	"context"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
)

type GetAllUsersQuery struct {
	userRepository repositories.UserRepository
}

func NewGetAllUsersQuery(userRepository repositories.UserRepository) GetAllUsersQuery {
	return GetAllUsersQuery{
		userRepository: userRepository,
	}
}

func (gu *GetAllUsersQuery) Do(ctx context.Context) ([]entities.User, error) {
	return gu.userRepository.GetAll(ctx)
}
