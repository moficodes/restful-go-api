package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/moficodes/restful-go-api/pkg/internal/handler"
	internalMiddleware "github.com/moficodes/restful-go-api/pkg/internal/middleware"
	"github.com/moficodes/restful-go-api/pkg/pkg/middleware"
)

func main() {
	s := &handler.Server{Routes: make([]string, 0)}
	r := mux.NewRouter()

	r.Handle("/", s)
	api := r.PathPrefix("/api/v1").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	auth.Use(internalMiddleware.JWTAuth)
	auth.HandleFunc("/check", handler.AuthHandler)

	api.Use(middleware.Time)

	api.HandleFunc("/users", handler.GetAllUsers).Methods(http.MethodGet)

	api.HandleFunc("/courses", handler.GetCoursesWithInstructorAndAttendee).
		Queries("instructor", "{instructor:[0-9]+}", "attendee", "{attendee:[0-9]+}").
		Methods(http.MethodGet)

	api.HandleFunc("/courses", handler.GetAllCourses).Methods(http.MethodGet)
	api.HandleFunc("/instructors", handler.GetAllInstructors).Methods(http.MethodGet)

	// in gorilla mux we can name path parameters
	// the library will put them in an key,val map for us
	api.HandleFunc("/users/{id}", handler.GetUserByID).Methods(http.MethodGet)
	api.HandleFunc("/courses/{id}", handler.GetCoursesByID).Methods(http.MethodGet)
	api.HandleFunc("/instructors/{id}", handler.GetInstructorByID).Methods(http.MethodGet)

	port := "7999"
	log.Println("starting web server on port", port)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		s.Routes = append(s.Routes, t)
		return nil
	})
	log.Println("available routes: ", s.Routes)
	// instead of using the default handler that comes with net/http we use the mux router from gorilla mux
	log.Fatal(http.ListenAndServe(":"+port, middleware.Chain(r, middleware.MuxLogger, middleware.Logger)))
}
