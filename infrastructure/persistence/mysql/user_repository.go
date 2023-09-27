package mysql

import (
	"context"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserRepository(ctx context.Context, db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	result := r.db.WithContext(r.ctx).Create(user)
	return result.Error
}
