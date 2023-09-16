package routes

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/controller"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes() *echo.Echo {
	e := echo.New()

	homeController := controller.NewHomeController()

	e.GET("/", homeController.Home)

	return e
}
