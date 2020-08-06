package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var (
	users       []User
	instructors []Instructor
	courses     []Course
)

type server struct {
	Routes []string `json:"routes"`
}

// User represent one user of our service
type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Interests []string `json:"interests"`
}

// Instructor type represent a instructor for a course
type Instructor struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Expertise []string `json:"expertise"`
}

// Course is course being taught
type Course struct {
	ID           int      `json:"id"`
	InstructorID int      `json:"instructor_id"`
	Name         string   `json:"name"`
	Topics       []string `json:"topics"`
	Attendees    []int    `json:"attendees"`
}

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
		set := make(map[int]User)
		for _, user := range users {

			if contains(user.Interests, interests) {
				set[user.ID] = user
			}
		}

		res := make([]User, len(set))
		idx := 0
		for _, v := range set {
			res[idx] = v
			idx++
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
		set := make(map[int]Instructor)
		for _, instructor := range instructors {
			if contains(instructor.Expertise, expertise) {
				set[instructor.ID] = instructor
			}
		}

		res := make([]Instructor, len(set))
		idx := 0
		for _, v := range set {
			res[idx] = v
			idx++
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
		set := make(map[int]Course)
		for _, course := range courses {
			if contains(course.Topics, topics) {
				set[course.ID] = course
			}
		}

		res := make([]Course, len(set))
		idx := 0
		for _, v := range set {
			res[idx] = v
			idx++
		}

		e := json.NewEncoder(w)
		e.Encode(res)
		return
	}

	e := json.NewEncoder(w)
	e.Encode(courses)
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

func main() {
	s := &server{Routes: make([]string, 0)}
	r := mux.NewRouter()
	r.Handle("/", s)

	r.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
	r.HandleFunc("/instructors", getAllInstructors).Methods(http.MethodGet)

	// in gorilla mux we can name path parameters
	// the library will put them in an key,val map for us
	r.HandleFunc("/users/{id}", getUserByID).Methods(http.MethodGet)
	r.HandleFunc("/courses/{id}", getUserByID).Methods(http.MethodGet)
	r.HandleFunc("/instructors/{id}", getUserByID).Methods(http.MethodGet)

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
	log.Fatal(http.ListenAndServe(":"+port, r))
}
