package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var courses []Course

func init() {
	loadCourses("./data/courses.json")
}

func loadCourses(path string) {
	if courses != nil {
		return
	}
	if err := readContent(path, &courses); err != nil {
		log.Println("Could not read courses data")
	}
}

func GetAllCourses(c echo.Context) error {
	topics := []string{}
	attendees := []string{}
	instructor := -1

	if err := echo.QueryParamsBinder(c).
		Strings("topic", &topics).
		Int("instructor", &instructor).
		Strings("attendee", &attendees).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	res := make([]Course, 0)
	for _, course := range courses {
		if contains(course.Topics, topics) && containsInt(course.Attendees, attendees) && (instructor == -1 || course.InstructorID == instructor) {
			res = append(res, course)
		}
	}
	return c.JSON(http.StatusOK, res)
}

func GetCoursesByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *Course
	for _, v := range courses {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "course with id not found")
	}

	return c.JSON(http.StatusCreated, data)
}
