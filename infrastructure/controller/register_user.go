package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
	"go-microservice/infrastructure/container"
	"net/http"
)

// UserRegistrationRequest represents the payload for registering a user.
type UserRegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterController struct{}

func NewRegisterController() RegisterController {
	return RegisterController{}
}

func (uc *RegisterController) Register(ctx echo.Context, c container.Container) error {
	var request UserRegistrationRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
	}

	createUserCommand := c.GetCreateUserByEmailCommand(context.Background())
	err := createUserCommand.Do(context.Background(), command.UserRegistrationRequest(request))
	if err != nil {
		return err
	}
	// TODO: Call the domain service to handle the actual registration logic and persist the user entity

	return ctx.JSON(http.StatusCreated, user)
}
