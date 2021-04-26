package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moficodes/restful-go-api/database/internal/datasource"
)

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.DB.GetAllUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	user, err := h.DB.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("user with id : %d not found", id))
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateNewUser(c echo.Context) error {
	c.Request().Header.Add("Content-Type", "application/json")
	if c.Request().ContentLength == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "body is required for this method")
	}
	user := new(datasource.User)
	err := c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "body could not be parsed")
	}

	id, err := h.DB.CreateNewUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not create user")
	}

	return c.JSON(http.StatusCreated, Message{Data: fmt.Sprintf("user created with id : %d", id)})
}

func (h *Handler) AddUserInterest(c echo.Context) error {
	c.Request().Header.Add("Content-Type", "application/json")

	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	if c.Request().ContentLength == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "body is required for this method")
	}

	var interests []string
	err := c.Bind(&interests)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "body not be valid")
	}
	count, err := h.DB.AddUserInterest(id, interests)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user interest")
	}
	return c.JSON(http.StatusCreated, Message{Data: fmt.Sprintf("%d user interest added", count)})
}
