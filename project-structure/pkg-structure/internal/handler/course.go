package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	courses []Course
)

func init() {
	if err := readContent("./data/courses.json", &courses); err != nil {
		log.Fatalln("Could not read courses data")
	}
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
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

func GetCoursesWithInstructorAndAttendee(w http.ResponseWriter, r *http.Request) {
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

func GetCoursesByID(w http.ResponseWriter, r *http.Request) {
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
