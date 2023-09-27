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

func NewCreateUserController() RegisterController {
	return RegisterController{}
}

func (uc *RegisterController) Create(ctx echo.Context, c container.Container) error {
	fmt.Println("test")
	var request UserRegistrationRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	fmt.Println("test2, ", request)
	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
	}

	fmt.Println("test3, ", request)
	createUserCommand := c.GetCreateUserByEmailCommand(context.Background())

	err := createUserCommand.Do(context.Background(), command.UserRegistrationRequest(request))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
