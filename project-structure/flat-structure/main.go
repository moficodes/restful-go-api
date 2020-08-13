package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	users       []User
	instructors []Instructor
	courses     []Course
)

func init() {
	if err := readContent("./data/courses.json", &courses); err != nil {
		log.Fatalln("Could not read courses data")
	}
	if err := readContent("./data/instructors.json", &instructors); err != nil {
		log.Fatalln("Could not read instructors data")
	}
	if err := readContent("./data/users.json", &users); err != nil {
		log.Fatalln("Could not read users data")
	}
}

func readContent(filename string, store interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, store)
}

func main() {
	s := &server{Routes: make([]string, 0)}
	r := mux.NewRouter()

	r.Handle("/", s)
	api := r.PathPrefix("/api/v1").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	auth.Use(JWTAuth)
	auth.HandleFunc("/check", authHandler)

	api.Use(Time)

	api.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)

	api.HandleFunc("/courses", getCoursesWithInstructorAndAttendee).
		Queries("instructor", "{instructor:[0-9]+}", "attendee", "{attendee:[0-9]+}").
		Methods(http.MethodGet)

	api.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
	api.HandleFunc("/instructors", getAllInstructors).Methods(http.MethodGet)

	// in gorilla mux we can name path parameters
	// the library will put them in an key,val map for us
	api.HandleFunc("/users/{id}", getUserByID).Methods(http.MethodGet)
	api.HandleFunc("/courses/{id}", getUserByID).Methods(http.MethodGet)
	api.HandleFunc("/instructors/{id}", getUserByID).Methods(http.MethodGet)

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
	log.Fatal(http.ListenAndServe(":"+port, Chain(r, MuxLogger, Logger)))
}
