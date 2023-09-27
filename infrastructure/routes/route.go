package routes

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller"
)

func SetupRoutes(c container.Container) *echo.Echo {
	e := echo.New()

	createUserController := controller.NewCreateUserController()
	retrieveUsersController := controller.NewRetrieveUsers()

	e.GET("/retrieve-users", func(ctx echo.Context) error {
		return retrieveUsersController.Retrieve(ctx, c)
	})

	e.POST("/register", func(ctx echo.Context) error {
		return createUserController.Create(ctx, c)
	})

	return e
}
