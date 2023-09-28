package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	domError "go-microservice/domain/error"
	"go-microservice/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const minPasswordLength = 8

type UserCreator interface {
	Create(credentials UserCreationRequest) error
}

type creator struct {
	ctx            context.Context
	userRepository repositories.UserRepository
	bcryptCost     int
}

func NewCreatorService(
	ctx context.Context,
	repository repositories.UserRepository,
	bcryptCost int,
) UserCreator {
	return &creator{
		ctx:            ctx,
		userRepository: repository,
		bcryptCost:     bcryptCost,
	}
}

type UserCreationRequest struct {
	Firstname string
	Surname   string
	Email     string
	Password  string
}

func (c *creator) Create(credentials UserCreationRequest) error {
	credentials.Email = strings.ToLower(credentials.Email)

	if len(credentials.Password) < minPasswordLength {
		return stacktrace.Propagate(
			domError.PasswordTooShort{MinLength: minPasswordLength},
			"error on password length: does not meet the minimum length requirement",
		)
	}

	existingUser, err := c.userRepository.GetUserByEmail(c.ctx, credentials.Email)
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on retrieving user with email: %s",
			credentials.Email,
		)
	}

	if existingUser != nil {
		return stacktrace.Propagate(
			domError.EmailAlreadyExists{Email: credentials.Email},
			"error on user creation: email is already in use: %s",
			credentials.Email,
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), c.bcryptCost)
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on hashing password for user: %s",
			credentials.Email,
		)
	}

	credentialsData := entities.User{
		UUID:      uuid.New(),
		Firstname: &credentials.Firstname,
		Surname:   &credentials.Surname,
		Email:     credentials.Email,
		Password:  string(hashedPassword),
	}

	err = c.userRepository.CreateUser(&credentialsData)
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on creating user with email: %s",
			credentials.Email,
		)
	}

	return nil
}
