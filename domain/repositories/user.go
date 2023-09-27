package repositories

import (
	"go-microservice/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
}
