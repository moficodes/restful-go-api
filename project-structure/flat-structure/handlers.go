package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(s.Routes)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	interests, ok := query["interest"]
	// the key was found.
	if ok {
		res := make([]User, 0)
		for _, user := range users {
			if contains(user.Interests, interests) {
				res = append(res, user)
			}
		}

		e := json.NewEncoder(w)
		e.Encode(res)
		return
	}

	e := json.NewEncoder(w)
	e.Encode(users)
}

func getAllInstructors(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	expertise, ok := query["expertise"]
	// the key was found.
	if ok {
		res := make([]Instructor, 0)
		for _, instructor := range instructors {
			if contains(instructor.Expertise, expertise) {
				res = append(res, instructor)
			}
		}

		e := json.NewEncoder(w)
		e.Encode(res)
		return
	}

	e := json.NewEncoder(w)
	e.Encode(instructors)
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	topics, ok := query["topic"]
	if ok {
		res := make([]Course, 0)
		for _, course := range courses {
			if contains(course.Topics, topics) {
				res = append(res, course)
			}
		}

		e := json.NewEncoder(w)
		e.Encode(res)
		return
	}

	e := json.NewEncoder(w)
	e.Encode(courses)
}

func getCoursesWithInstructorAndAttendee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// we don't have to check for multiple instructor because the way our data is structured
	// there is no way multiple instructor can be part of same course
	_instructor := r.URL.Query().Get("instructor")
	instructorID, _ := strconv.Atoi(_instructor)
	// but multiple attendee can be part of the same course
	// since we gurrantee only valid integer queries will be sent to this route
	// we don't need to check if there is value or not.
	attendees := r.URL.Query()["attendee"]
	res := make([]Course, 0)

	for _, course := range courses {
		if course.InstructorID == instructorID && containsInt(course.Attendees, attendees) {
			res = append(res, course)
		}
	}

	e := json.NewEncoder(w)
	e.Encode(res)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	pathParams := mux.Vars(r)
	id := -1
	var err error
	if val, ok := pathParams["id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "need a valid id"}`))
			return
		}
	}

	var data *User
	for _, v := range users {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	}

	e := json.NewEncoder(w)
	e.Encode(data)
}

func getCoursesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// this function takes in the request and parses
	// all the pathParams from it
	// pathParams is a map[string]string
	pathParams := mux.Vars(r)
	id := -1
	var err error
	if val, ok := pathParams["id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "need a valid id"}`))
			return
		}
	}

	var data *Course
	for _, v := range courses {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	}

	e := json.NewEncoder(w)
	e.Encode(data)
}

func getInstructorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	pathParams := mux.Vars(r)
	id := -1
	var err error
	if val, ok := pathParams["id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "need a valid id"}`))
			return
		}
	}

	var data *Instructor
	for _, v := range instructors {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	}

	e := json.NewEncoder(w)
	e.Encode(data)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(ContextKey("props")).(jwt.MapClaims)

	name, ok := data["name"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "not authorized"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"message\": \"hello %v\"}", name)))
}

func contains(in []string, val []string) bool {
	found := 0

	for _, n := range in {
		n = strings.ToLower(n)
		for _, v := range val {
			if n == strings.ToLower(v) {
				found++
				break
			}
		}
	}

	return len(val) == found
}

func containsInt(in []int, val []string) bool {
	found := 0
	for _, _n := range in {
		n := strconv.Itoa(_n)
		for _, v := range val {
			if n == v {
				found++
				break
			}
		}
	}

	return len(val) == found
}
