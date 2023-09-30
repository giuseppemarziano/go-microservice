package service

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/palantir/stacktrace"
	domError "go-microservice/domain/error"
	"go-microservice/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserAuthenticator interface {
	Authenticate(credentials UserAuthenticationRequest) (*string, error)
}

type userAuthenticationService struct {
	ctx            context.Context
	userRepository repositories.UserRepository
}

func NewUserAuthenticationService(
	ctx context.Context,
	repository repositories.UserRepository,
) UserAuthenticator {
	return &userAuthenticationService{
		ctx:            ctx,
		userRepository: repository,
	}
}

type UserAuthenticationRequest struct {
	Email    string
	Password string
}

func (s *userAuthenticationService) Authenticate(credentials UserAuthenticationRequest) (*string, error) {
	user, err := s.userRepository.GetUserByEmail(s.ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, domError.UserNotFound{
			Email: credentials.Email,
		}) {
			return nil, stacktrace.Propagate(
				err,
				"User not found")
		}
		return nil, stacktrace.Propagate(
			err,
			"error on retrieving user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return nil, stacktrace.Propagate(
			err,
			"invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.UUID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, stacktrace.Propagate(
			err,
			"error on generating user token",
		)
	}

	return &tokenString, nil
}
