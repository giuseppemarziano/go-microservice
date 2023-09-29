package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"net/http"
)

type GetUserByEmail struct{}

func NewGetUserByEmail() GetUserByEmail {
	return GetUserByEmail{}
}

func (gc GetUserByEmail) RetrieveByEmail(echo echo.Context, c container.Container) error {
	ctx := context.Background()

	email := echo.Param("email")

	getByEmailQuery := c.GetGetUserByEmailQuery(ctx)
	users, err := getByEmailQuery.Do(ctx, email)
	if err != nil {
		return echo.JSON(
			http.StatusBadRequest,
			nil,
		)
	}

	return echo.JSON(http.StatusOK, users)
}
