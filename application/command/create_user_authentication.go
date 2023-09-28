package command

import (
	"context"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/service"
)

type AuthenticateUserCommand struct {
	authenticationService service.UserAuthenticator
}

func NewAuthenticateUserCommand(authenticationService service.UserAuthenticator) AuthenticateUserCommand {
	return AuthenticateUserCommand{
		authenticationService: authenticationService,
	}
}

type UserAuthenticationRequest struct {
	Email    string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *AuthenticateUserCommand) Do(ctx context.Context, credentials UserAuthenticationRequest) (*string, error) {
	serviceCredentials := service.UserAuthenticationRequest{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	token, err := c.authenticationService.Authenticate(serviceCredentials)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to authenticate user")
	}

	return token, nil
}
