package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moficodes/restful-go-api/restapi-docs-echo/pkg/middleware"
)

// Authenticated godoc
// @Summary Get user info from the JWT token
// @Description get user info from the JWT token
// @Tags authenticated
// @Accept */*
// @Produce json
// @Success 200 {object} Message
// @Failure 401 {object} map[string]interface{}
// @Router /auth/test [get]
func Authenticated(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	_name, ok := cc.Claims["name"]
	if !ok {
		echo.NewHTTPError(http.StatusUnauthorized, "malformed jwt")
	}

	name := fmt.Sprintf("%v", _name)

	return c.JSON(http.StatusOK, Message{Data: name})
}
