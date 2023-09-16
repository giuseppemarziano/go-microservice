package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() HomeController {
	return HomeController{}
}

func (hc *HomeController) Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}
