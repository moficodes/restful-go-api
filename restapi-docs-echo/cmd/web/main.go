package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/restful-go-api/restapi-docs-echo/internal/handler"
	"github.com/moficodes/restful-go-api/restapi-docs-echo/pkg/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/moficodes/restful-go-api/restapi-docs-echo/docs"
)

type server struct{}

func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

// @title O'Reilly RESTful Go API Course
// @version 1.0
// @description This is a demo of openapi spec generation with Go.
// @termsOfService http://swagger.io/terms/

// @contact.name Mofi Rahman
// @contact.url http://www.swagger.io/support
// @contact.email moficodes@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7999
// @BasePath /
// @schemes http
func main() {
	e := echo.New()

	specialLogger := echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, \n",
	})
	e.Use(middleware.Logger, specialLogger)
	e.Use(echoMiddleware.CORS())
	e.Use(echoMiddleware.Recover())

	e.GET("/", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
