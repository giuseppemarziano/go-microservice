package route

import (
	"github.com/labstack/echo/v4"
	"main/infrastructure/controller" // adjust the import path
)

// NewRouter initializes and returns a new router
func Init(e *echo.Echo) {
	e.GET("/hello", controller.HelloController)
}
