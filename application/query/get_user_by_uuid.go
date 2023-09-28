package query

import (
	"context"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
)

type GetUserByUUIDQuery struct {
	userRepository repositories.UserRepository
}

func NewGetUserByUUIDQuery(userRepository repositories.UserRepository) GetUserByUUIDQuery {
	return GetUserByUUIDQuery{
		userRepository: userRepository,
	}
}

func (guu *GetUserByUUIDQuery) Do(ctx context.Context, uuid string) (*entities.User, error) {
	user, err := guu.userRepository.GetUserByUUID(ctx, uuid)
	if err != nil {
		return nil, stacktrace.Propagate(err, "error retrieving user by uuid: %s", uuid)
	}
	return user, nil
}
