package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCourses(c echo.Context) error {
	topics := []string{}
	attendees := []string{}
	instructor := -1

	if err := echo.QueryParamsBinder(c).
		Strings("topic", &topics).
		Int("instructor", &instructor).
		Strings("attendee", &attendees).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	courses, err := h.DB.GetAllCourses()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}
	return c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCoursesByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	course, err := h.DB.GetCoursesByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	if course == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("course with id : %d not found", id))
	}

	return c.JSON(http.StatusOK, course)
}

func (h *Handler) GetCoursesForInstructor(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("instructorID", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}
	courses, err := h.DB.GetCoursesForInstructor(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}
	return c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCoursesForUser(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("userID", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}
	courses, err := h.DB.GetCoursesForUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}
	return c.JSON(http.StatusOK, courses)
}
