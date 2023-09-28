package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-microservice/application/command"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller/response"
	"net/http"
)

type CreateUserController struct {
}

type UserCreationRequest struct { // Define the UserRegistrationRequest type
	Firstname string `json:"firstname" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewCreateUserController() CreateUserController {
	return CreateUserController{}
}

func (rc CreateUserController) Create(echo echo.Context, c container.Container) error {
	var request UserCreationRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, response.NewCreateUserResponse("invalid request payload"))
	}

	if err := echo.Validate(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, response.NewCreateUserResponse("error on validation"))
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
		return echo.JSON(
			http.StatusInternalServerError,
			response.NewCreateUserResponse("internal server error"),
		)
	}

	return echo.JSON(http.StatusCreated, "User created successfully")
}
