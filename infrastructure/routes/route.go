package routes

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/controller"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()

	homeController := controller.NewHomeController()

	e.GET("/", homeController.Home)

	return e
}
