// domain/repositories/user_repository.go
package repositories

import (
	"context"
	"go-microservice/domain/entities"
	"go-microservice/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) SaveUser(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
