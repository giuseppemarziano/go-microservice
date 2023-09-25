package controller

import (
	"github.com/labstack/echo/v4"
	"go-microservice/domain/entities"
	"net/http"
)

type RegisterController struct{}

func NewRegisterController() RegisterController {
	return RegisterController{}
}

func (uc *RegisterController) Register(c echo.Context) error {
	var request entities.User
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request") // TODO fix error message
	}

	return c.JSON(http.StatusCreated, user)
}
