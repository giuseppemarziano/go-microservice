package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"net/http"
)

type GetUserByUUID struct{}

func NewGetUserByUUID() GetUserByUUID {
	return GetUserByUUID{}
}

func (gc GetUserByUUID) RetrieveByUUID(echo echo.Context, c container.Container) error {
	ctx := context.Background()

	uuid := echo.Param("uuid")

	getByUUIDQuery := c.GetGetUserByUUIDQuery(ctx)
	users, err := getByUUIDQuery.Do(ctx, uuid)
	if err != nil {
		return echo.JSON(
			http.StatusBadRequest,
			nil,
		)
	}

	return echo.JSON(http.StatusOK, users)
}
