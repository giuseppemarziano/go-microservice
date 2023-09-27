package routes

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"go-microservice/infrastructure/controller"
)

func SetupRoutes(c container.Container) *echo.Echo {
	e := echo.New()

	homeController := controller.NewHomeController()
	registerController := controller.NewRegisterController()

	e.GET("/", homeController.Home)

	e.POST("/register", func(ctx echo.Context) error {
		return registerController.Register(ctx, c)
	})

	return e
}
