package controller

import (
	"context"
	"fmt"
	"go-microservice/infrastructure/container"
	"net/http"

	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
)

// RegisterController is an empty struct as per requirement.
type RegisterController struct{}

type UserRegistrationRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewCreateUserController() RegisterController {
	return RegisterController{}
}

func (rc RegisterController) Create(echo echo.Context, c container.Container) error {
	var request command.UserRegistrationRequest
	if err := echo.Bind(&request); err != nil {
		fmt.Println(err)
		return echo.JSON(http.StatusBadRequest, "error")
	}

	ctx := context.Background()

	createUserCommand := c.GetCreateUserByEmailCommand(ctx)

	err := createUserCommand.Do(ctx, request)

	if err != nil {
		fmt.Println("Failed to create user:", err)
		return echo.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return echo.JSON(http.StatusCreated, "User created successfully")
}
