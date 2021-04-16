package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	users       []User
	instructors []Instructor
	courses     []Course
)

type server struct{}

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

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
// 	e := json.NewEncoder(w)
// 	e.Encode(s.Routes)
// }

func getAllUsers(c echo.Context) error {
	interests := []string{}
	if err := echo.QueryParamsBinder(c).Strings("interest", &interests).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	res := make([]User, 0)
	for _, user := range users {
		if contains(user.Interests, interests) {
			res = append(res, user)
		}
	}

	return c.JSON(http.StatusOK, res)
}

func getAllInstructors(c echo.Context) error {
	expertise := []string{}

	// the key was found.
	if err := echo.QueryParamsBinder(c).Strings("expertise", &expertise).BindError(); err != nil { //watch the == here
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}
	res := make([]Instructor, 0)
	for _, instructor := range instructors {
		if contains(instructor.Expertise, expertise) {
			res = append(res, instructor)
		}
	}
	return c.JSON(http.StatusOK, res)
}

func getAllCourses(c echo.Context) error {
	topics := []string{}
	attendees := []string{}
	instructor := -1

	if err := echo.QueryParamsBinder(c).
		Strings("topic", &topics).
		Int("instructor", &instructor).
		Strings("attendee", &attendees).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")
	}

	res := make([]Course, 0)
	for _, course := range courses {
		if contains(course.Topics, topics) && containsInt(course.Attendees, attendees) && (instructor == -1 || course.InstructorID == instructor) {
			res = append(res, course)
		}
	}
	return c.JSON(http.StatusOK, res)
}

func getUserByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *User
	for _, v := range users {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user with id not found")
	}

	return c.JSON(http.StatusOK, data)
}

func getCoursesByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *Course
	for _, v := range courses {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user with id not found")
	}

	return c.JSON(http.StatusOK, data)
}

func getInstructorByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	var data *Instructor
	for _, v := range instructors {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user with id not found")
	}

	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()
	e.GET("/users", getAllUsers)
	e.GET("/instructors", getAllInstructors)
	e.GET("/courses", getAllCourses)

	e.GET("/users/:id", getUserByID)
	e.GET("/instructors/:id", getInstructorByID)
	e.GET("/courses/:id", getCoursesByID)
	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))
}
