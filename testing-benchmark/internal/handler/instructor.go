package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	instructors []Instructor
)

func init() {
	loadInstructors("./data/instructors.json")
}

func loadInstructors(path string) {
	if err := readContent(path, &instructors); err != nil {
		log.Println("Could not read instructors data")
	}
}

func GetAllInstructors(w http.ResponseWriter, r *http.Request) {
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

func GetInstructorByID(w http.ResponseWriter, r *http.Request) {
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
