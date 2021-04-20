package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var users []User

func init() {
	loadUsers("./data/users.json")
}

func loadUsers(path string) {
	if err := readContent(path, &users); err != nil {
		log.Println("Could not read instructors data")
	}
}

func GetAllUsers(c echo.Context) error {
	interests := []string{}
	if err := echo.QueryParamsBinder(c).Strings("interest", &interests).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	res := make([]User, 0)
	for _, user := range users {
		if contains(user.Interests, interests) {
			res = append(res, user)
		}
	}

	return c.JSON(http.StatusOK, res)
}

func GetUserByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *User
	for _, v := range users {
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
