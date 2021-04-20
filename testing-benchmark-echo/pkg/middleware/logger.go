package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(c.Request().URL)
		return next(c)
	}
}
