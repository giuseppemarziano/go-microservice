package routes

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/controller"
	"go-microservice/infrastructure/http"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes() *echo.Echo {
	e := echo.New()

	homeController := controller.NewHomeController()

	e.GET("/", homeController.Home, http.AuthenticationMiddleware)

	return e
}
