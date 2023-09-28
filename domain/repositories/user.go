package repositories

import (
	"context"
	"go-microservice/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetAll(ctx context.Context) ([]entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (*entities.User, error)
}
