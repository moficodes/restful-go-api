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

// API godoc
// @Summary Get all users
// @Description get all users matching given query params. returns all by default.
// @Tags API
// @Accept */*
// @Produce json
// @Param interest query string false "interests to filter by"
// @Success 200 {object} []User
// @Router /api/v1/users [get]
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

// API godoc
// @Summary Get user by id
// @Description get user matching given ID.
// @Tags API
// @Accept */*
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} User
// @Failure 404 {object} HTTPError
// @Router /api/v1/users/{id} [get]
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
