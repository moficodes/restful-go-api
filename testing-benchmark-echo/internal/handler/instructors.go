package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var instructors []Instructor

func init() {
	loadInstructors("./data/instructors.json")
}

func loadInstructors(path string) {
	if err := readContent(path, &instructors); err != nil {
		log.Println("Could not read instructors data")
	}
}

func GetAllInstructors(c echo.Context) error {
	expertise := []string{}

	// the key was found.
	if err := echo.QueryParamsBinder(c).Strings("expertise", &expertise).BindError(); err != nil { //watch the == here
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}
	res := make([]Instructor, 0)
	for _, instructor := range instructors {
		if contains(instructor.Expertise, expertise) {
			res = append(res, instructor)
		}
	}
	return c.JSON(http.StatusOK, res)
}

func GetInstructorByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *Instructor
	for _, v := range instructors {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user with id not found")
	}

	return c.JSON(http.StatusOK, data)
}
