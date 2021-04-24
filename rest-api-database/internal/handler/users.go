package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllUsers(c echo.Context) error {
	interests := []string{}
	if err := echo.QueryParamsBinder(c).Strings("interest", &interests).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

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

	return c.JSON(http.StatusOK, user)
}
