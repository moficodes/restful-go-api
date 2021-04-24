package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllInstructors(c echo.Context) error {
	expertise := []string{}

	// the key was found.
	if err := echo.QueryParamsBinder(c).Strings("expertise", &expertise).BindError(); err != nil { //watch the == here
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	instructors, err := h.DB.GetAllInstructors()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	return c.JSON(http.StatusOK, instructors)
}

func (h *Handler) GetInstructorByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	instructor, err := h.DB.GetInstructorByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	return c.JSON(http.StatusOK, instructor)
}
