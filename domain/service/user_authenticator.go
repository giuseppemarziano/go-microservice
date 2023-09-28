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
	Authenticate(credentials UserAuthenticationRequest) (string, error)
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

func (s *userAuthenticationService) Authenticate(credentials UserAuthenticationRequest) (string, error) {
	user, err := s.userRepository.GetUserByEmail(s.ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, domError.UserNotFound{
			Email: credentials.Email,
		}) {
			return "", stacktrace.Propagate(err, "User not found")
		}
		return "", stacktrace.Propagate(err, "Failed to retrieve user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", stacktrace.Propagate(err, "Invalid credentials")
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.UUID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", stacktrace.Propagate(err, "Failed to generate user token")
	}

	return tokenString, nil
}
