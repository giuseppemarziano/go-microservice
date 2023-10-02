package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller"
	"go-microservice/infrastructure/http"
)

func Routes(c container.Container) *echo.Echo {
	e := echo.New()

	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	createUserController := controller.NewCreateUserController()
	retrieveUsersController := controller.NewRetrieveUsers()
	retrieveUserByUUIDController := controller.NewGetUserByUUID()
	retrieveUserByEmailController := controller.NewGetUserByEmail()
	authController := controller.NewAuthController()

	e.GET(
		"/retrieve-users",
		func(ctx echo.Context) error {
			return retrieveUsersController.RetrieveAll(ctx, c)
		},
		http.AuthenticationMiddleware,
	)

	e.POST(
		"/get-user-by-uuid",
		func(ctx echo.Context) error {
			return retrieveUserByUUIDController.RetrieveByUUID(ctx, c)
		},
	)

	e.POST(
		"/register",
		func(ctx echo.Context) error {
			return createUserController.Create(ctx, c)
		},
	)

	e.POST(
		"/get-user-by-uuid",
		func(ctx echo.Context) error {
			return retrieveUserByEmailController.RetrieveByEmail(ctx, c)
		},
	)

	e.POST(
		"/login",
		func(ctx echo.Context) error {
			return authController.Login(ctx, c)
		},
	)

	return e
}
