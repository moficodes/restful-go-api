package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/restful-go-api/database/internal/datasource"
	"github.com/moficodes/restful-go-api/database/internal/handler"
	"github.com/moficodes/restful-go-api/database/pkg/database"
	"github.com/moficodes/restful-go-api/database/pkg/middleware"
)

type server struct{}

func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func main() {
	connStr := "postgresql://postgres:password@localhost:5432/postgres"
	pool, err := database.PGPool(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Close()

	p := datasource.NewPostgres(pool)
	h := handler.NewHandler(p)

	e := echo.New()

	specialLogger := echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, \n",
	})
	e.Use(middleware.Logger, specialLogger)

	auth := e.Group("/auth")
	auth.Use(middleware.JWT)
	auth.GET("/test", handler.Authenticated)
	api := e.Group("/api/v1")
	_ = Chain(h.GetAllUsers, middleware.Logger, specialLogger) // this would give us a new handler that we can use in place of any other handler
	api.GET("/users", h.GetAllUsers)
	api.GET("/instructors", h.GetAllInstructors)
	api.GET("/courses", h.GetAllCourses)

	api.GET("/users/:id", h.GetUserByID)
	api.GET("/instructors/:id", h.GetInstructorByID)
	api.GET("/courses/:id", h.GetCoursesByID)
	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))
}
