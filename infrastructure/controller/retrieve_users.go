package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-microservice/infrastructure/container"
	"net/http"
)

type RetrieveUsers struct{}

func NewRetrieveUsers() RetrieveUsers {
	return RetrieveUsers{}
}

func (gu *RetrieveUsers) Retrieve(ctx echo.Context, c container.Container) error {
	getAllUsersQuery := c.GetGetAllUsersQuery(ctx.Request().Context())

	users, err := getAllUsersQuery.Do(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to retrieve users: " + err.Error()})
	}

	return ctx.JSON(http.StatusOK, users)
}
