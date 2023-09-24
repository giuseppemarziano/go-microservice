package repositories

import (
	"context"
	"go-microservice/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
}
