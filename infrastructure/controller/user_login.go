package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
	"go-microservice/infrastructure/container"
	"net/http"
)

type AuthController struct {
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (ac AuthController) Login(echo echo.Context, c container.Container) error {
	var request LoginRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, "error")
	}

	if err := echo.Validate(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, "error")
	}

	ctx := context.Background()

	authenticateUserCommand := c.GetCreateUserAuthenticationCommand(ctx)
	token, err := authenticateUserCommand.Do(ctx, command.UserAuthenticationRequest{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return echo.JSON(http.StatusUnauthorized, "invalid credentials")
	}

	return echo.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
