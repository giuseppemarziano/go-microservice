package http

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "authentication failed")
		}

		splitToken := strings.Split(authorization, "Bearer ")
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "authentication failed")
		}
		tokenString := splitToken[1]

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			fmt.Println("Secret key not set")
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "authentication failed")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "authentication failed")
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return echo.NewHTTPError(http.StatusUnauthorized, "token expired")
			}
		}

		if nbf, ok := claims["nbf"].(float64); ok {
			if time.Unix(int64(nbf), 0).After(time.Now()) {
				return echo.NewHTTPError(http.StatusUnauthorized, "token not valid yet")
			}
		}

		return next(c)
	}
}
