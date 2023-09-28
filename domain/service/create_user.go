package service

import (
	"context"
	"errors"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	domError "go-microservice/domain/error"
	"go-microservice/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
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

func NewCreatorService(ctx context.Context, repository repositories.UserRepository, bcryptCost int) UserCreator {
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

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(credentials.Email) {
		return stacktrace.Propagate(errors.New("invalid email format"), "error on validating email format: %s", credentials.Email)
	}

	if len(credentials.Password) < minPasswordLength {
		return stacktrace.Propagate(domError.PasswordTooShort{MinLength: minPasswordLength}, "error on password length: does not meet the minimum length requirement")
	}

	existingUser, err := c.userRepository.GetUserByEmail(c.ctx, credentials.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return stacktrace.Propagate(err, "error on retrieving user with email: %s", credentials.Email)
	}

	if existingUser != nil {
		return stacktrace.Propagate(domError.EmailAlreadyExists{Email: credentials.Email}, "error on user creation: email is already in use: %s", credentials.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), c.bcryptCost)
	if err != nil {
		return stacktrace.Propagate(err, "error on hashing password for user: %s", credentials.Email)
	}

	credentialsData := entities.User{
		Firstname: &credentials.Firstname,
		Surname:   &credentials.Surname,
		Email:     credentials.Email,
		Password:  string(hashedPassword),
	}

	err = c.userRepository.CreateUser(&credentialsData)
	if err != nil {
		return stacktrace.Propagate(err, "error on creating user with email: %s", credentials.Email)
	}

	return nil
}
