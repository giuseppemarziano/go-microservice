package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
	"go-microservice/infrastructure/container"
	"net/http"
)

type CreateUserController struct{}

type UserRegistrationRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewCreateUserController() CreateUserController {
	return CreateUserController{}
}

func (rc CreateUserController) Create(echo echo.Context, c container.Container) error {
	var request UserRegistrationRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := echo.Validate(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err))
	}

	ctx := context.Background()

	createUserCommand := c.GetCreateUserByEmailCommand(ctx)

	err := createUserCommand.Do(ctx, command.UserRegistrationRequest{
		Firstname: request.Firstname,
		Surname:   request.Surname,
		Email:     request.Email,
		Password:  request.Password,
	})

	if err != nil {
		fmt.Println("Failed to create user:", err)
		return echo.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return echo.JSON(http.StatusCreated, "User created successfully")
}
