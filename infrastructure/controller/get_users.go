package controller

import (
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"net/http"
)

type RetrieveUsers struct{}

func NewRetrieveUsers() RetrieveUsers {
	return RetrieveUsers{}
}

func (gu *RetrieveUsers) RetrieveAll(ctx echo.Context, c container.Container) error {
	getAllUsersQuery := c.GetGetAllUsersQuery(ctx.Request().Context())
	users, err := getAllUsersQuery.Do(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			nil,
		)
	}

	return ctx.JSON(http.StatusOK, users)
}
