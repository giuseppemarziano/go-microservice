package mysql

import (
	"context"
	"errors"
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

func (r *userRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user with the given email does not exist")
		}
		return nil, result.Error
	}
	return &user, nil
}
