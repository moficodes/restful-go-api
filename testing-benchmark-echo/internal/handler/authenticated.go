package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moficodes/restful-go-api/echo-testing/pkg/middleware"
)

func Authenticated(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	_name, ok := cc.Claims["name"]
	if !ok {
		echo.NewHTTPError(http.StatusUnauthorized, "malformed jwt")
	}

	name := fmt.Sprintf("%v", _name)

	return c.JSON(http.StatusOK, Message{Data: name})
}
