package controller

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/palantir/stacktrace"
	"go-microservice/application/command"
	domError "go-microservice/domain/error"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller/response"
	"log"
	"net/http"
)

type CreateUserController struct{}

type UserCreationRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewCreateUserController() CreateUserController {
	return CreateUserController{}
}

func (rc CreateUserController) Create(echoContext echo.Context, c container.Container) error {
	var request UserCreationRequest
	if err := echoContext.Bind(&request); err != nil {
		return echoContext.JSON(
			http.StatusBadRequest,
			response.NewCreateUserResponse("error on validating request payload"),
		)
	}

	if err := echoContext.Validate(&request); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "error on validation")
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
		rootErr := stacktrace.RootCause(err)

		var emailErr domError.EmailAlreadyExists
		var passwordErr domError.PasswordTooShort

		switch {
		case errors.As(rootErr, &emailErr):
			return echoContext.JSON(
				http.StatusConflict,
				response.NewCreateUserResponse(emailErr.Error()),
			)
		case errors.As(rootErr, &passwordErr):
			return echoContext.JSON(
				http.StatusBadRequest,
				response.NewCreateUserResponse(passwordErr.Error()),
			)
		default:
			log.Printf("Error creating user: %+v\n", err)

			return echoContext.JSON(
				http.StatusInternalServerError,
				response.NewCreateUserResponse("Failed to create user"),
			)
		}
	}

	return echoContext.JSON(http.StatusCreated, "User created successfully")
}
