package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type server struct {
	// becasue User is a property of server struct
	// we have access to user at all the server instance
	User User
}

type StatusOK struct {
	Message string `json:"message"`
}

type User struct {
	// this tags after our struct lets us change the represantation in json
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func (s *server) home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func (s *server) getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, s.User)
}

func (s *server) updateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	log.Println(u)

	s.User = *u
	return c.JSON(http.StatusOK, StatusOK{Message: "update ok"})
}

func main() {
	s := &server{
		User: User{
			Username: "moficodes",
			Email:    "moficodes@gmail.com",
			Age:      27,
		},
	}

	e := echo.New()
	e.GET("/", s.home)
	e.GET("/user", s.getUser)
	e.PUT("/user", s.updateUser)

	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))
}
