package mysql

import (
	"context"
	"errors"
	"github.com/palantir/stacktrace"
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
	if result.Error != nil {
		return stacktrace.Propagate(
			result.Error,
			"error on creating user with email: %s",
			user.Email,
		)
	}
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, stacktrace.Propagate(
			result.Error,
			"error on retrieving all users",
		)
	}
	return users, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, stacktrace.Propagate(
				result.Error,
				"error on retrieving user: user with email %s does not exist",
				email,
			)
		}
		return nil, stacktrace.Propagate(
			result.Error,
			"error on retrieving user with email: %s",
			email,
		)
	}
	return &user, nil
}

func (r *userRepository) GetUserByUUID(ctx context.Context, uuid string) (*entities.User, error) {
	var user entities.User
	result := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, stacktrace.Propagate(
				result.Error,
				"error on retrieving user: user with UUID %s does not exist",
				uuid,
			)
		}
		return nil, stacktrace.Propagate(
			result.Error,
			"error on retrieving user with UUID: %s",
			uuid,
		)
	}
	return &user, nil
}
