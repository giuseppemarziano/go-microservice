package repositories

import (
	"context"
	"go-microservice/domain/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
}
