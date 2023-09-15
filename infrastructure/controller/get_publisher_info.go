package controller

import "net/http"

func HelloController(c echo.Context) error {
	cc := c.(*context.CustomContext)
	// You can access services like cc.Services.Logger or cc.Services.Database here

	return c.String(http.StatusOK, "Hello, World!")
}