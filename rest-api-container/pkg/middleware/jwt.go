package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Claims jwt.MapClaims
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		// auth token have the structure `bearer <token>`
		// so we split it on the ` ` (space character)
		splitToken := strings.Split(authorization, " ")
		// if we end up with a array of size 2 we have the token as the
		// 2nd item in the array
		if len(splitToken) != 2 {
			// we got something different
			return echo.NewHTTPError(http.StatusUnauthorized, "no valid token found")
		}
		// second item is our possible token
		jwtToken := splitToken[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("very-secret"), nil
		})

		if err != nil {
			// we got something different
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			cc := &CustomContext{c, claims}
			return next(cc)

		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
	}
}
