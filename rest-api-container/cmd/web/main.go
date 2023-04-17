package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/restful-go-api/database/internal/datasource"
	"github.com/moficodes/restful-go-api/database/internal/handler"
	"github.com/moficodes/restful-go-api/database/pkg/middleware"

	_ "github.com/joho/godotenv/autoload"
)

type server struct{}

func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

// getDB creates a connection to the database
// based on environment variables.
func getDB() (*sql.DB, func() error) {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
		log.Fatalf("Error on pgxv4.RegisterDriver: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", os.Getenv("INSTANCE_CONNECTION_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	db, err := sql.Open("cloudsql-postgres", dsn)
	if err != nil {
		log.Fatalf("Error on sql.Open: %v", err)
	}

	return db, cleanup
}

func main() {
	// var (
	// 	dbUser    = os.Getenv("DB_USER")       // e.g. 'my-db-user'
	// 	dbPwd     = os.Getenv("DB_PASS")       // e.g. 'my-db-password'
	// 	dbTCPHost = os.Getenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	// 	dbPort    = os.Getenv("DB_PORT")       // e.g. '5432'
	// 	dbName    = os.Getenv("DB_NAME")       // e.g. 'my-database'
	// )

	// dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
	// 	dbTCPHost, dbUser, dbPwd, dbPort, dbName)
	// pool, err := database.PGPool(context.Background(), dbURI)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer pool.Close()

	db, cleanup := getDB()
	defer cleanup()

	p := datasource.NewSQL(db)
	// p := datasource.NewPostgres(pool)
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
	api.GET("/users", h.GetAllUsers)
	api.GET("/instructors", h.GetAllInstructors)
	api.GET("/courses", h.GetAllCourses)

	api.GET("/users/:id", h.GetUserByID)
	api.GET("/instructors/:id", h.GetInstructorByID)
	api.GET("/courses/:id", h.GetCoursesByID)

	api.GET("/courses/instructor/:instructorID", h.GetCoursesForInstructor)
	api.GET("/courses/user/:userID", h.GetCoursesForUser)

	api.POST("/users", h.CreateNewUser)
	api.POST("/users/:id/interests", h.AddUserInterest)

	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))
}
