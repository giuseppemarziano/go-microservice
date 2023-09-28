package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller"
	"go-microservice/infrastructure/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func SetupRoutes(c container.Container) *echo.Echo {
	e := echo.New()

	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	createUserController := controller.NewCreateUserController()
	retrieveUsersController := controller.NewRetrieveUsers()
	authController := controller.NewAuthController()

	e.GET(
		"/retrieve-users",
		func(ctx echo.Context) error {
			return retrieveUsersController.Retrieve(ctx, c)
		},
		http.AuthenticationMiddleware,
	)

	e.POST(
		"/register",
		func(ctx echo.Context) error {
			return createUserController.Create(ctx, c)
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
