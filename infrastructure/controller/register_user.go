package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
	"go-microservice/infrastructure/container"
	"net/http"
)

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
		fmt.Println("Error")
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	fmt.Println("Hello world.")

	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
	}

	createUserCommand := c.GetCreateUserByEmailCommand(context.Background())

	err := createUserCommand.Do(context.Background(), command.UserRegistrationRequest(request))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
