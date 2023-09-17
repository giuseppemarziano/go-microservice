package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		splitToken := strings.Split(authorization, "Bearer ")
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
		}
		tokenString := splitToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("YourSecretKey"), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		return next(c)
	}
}
