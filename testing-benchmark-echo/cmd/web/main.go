package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/restful-go-api/echo-testing/internal/handler"
	"github.com/moficodes/restful-go-api/echo-testing/pkg/middleware"
)

type server struct{}

func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func main() {
	e := echo.New()

	specialLogger := echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, \n",
	})
	e.Use(middleware.Logger, specialLogger)

	auth := e.Group("/auth")
	auth.Use(middleware.JWT)
	auth.GET("/test", handler.Authenticated)
	api := e.Group("/api/v1")
	_ = Chain(handler.GetAllUsers, middleware.Logger, specialLogger) // this would give us a new handler that we can use in place of any other handler
	api.GET("/users", handler.GetAllUsers)
	api.GET("/instructors", handler.GetAllInstructors)
	api.GET("/courses", handler.GetAllCourses)

	api.GET("/users/:id", handler.GetUserByID)
	api.GET("/instructors/:id", handler.GetInstructorByID)
	api.GET("/courses/:id", handler.GetCoursesByID)
	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))
}
