package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/moficodes/restful-go-api/database/internal/datasource"
)

type mock struct {
	users       []datasource.User
	courses     []datasource.Course
	instructors []datasource.Instructor
}

var (
	course = datasource.Course{
		ID:           1,
		InstructorID: 1,
		Name:         "Test Course",
		Topics:       []string{"go"},
		Attendees:    []int{1},
	}
)

func TestHandler_GetAllCourses(t *testing.T) {
	m := &mock{courses: []datasource.Course{course}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/courses", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	if assert.NoError(t, h.GetAllCourses(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var courses []datasource.Course
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &courses))
		assert.Equal(t, course, courses[0])
	}
}

func TestHandler_GetCoursesByID_success(t *testing.T) {
	m := &mock{courses: []datasource.Course{course}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/courses/1", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/courses/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetCoursesByID(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var course *datasource.Course
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &course))
		assert.Equal(t, 1, course.ID)
	}
}

func TestHandler_GetCoursesByID_failure(t *testing.T) {
	m := &mock{courses: []datasource.Course{course}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/courses/5", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/courses/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")

	assert.Error(t, h.GetUserByID(c), "should return error")
}

func (m *mock) GetAllCourses() ([]datasource.Course, error) {
	return m.courses, nil
}

func (m *mock) GetAllUsers() ([]datasource.User, error) {
	return nil, nil
}

func (m *mock) GetAllInstructors() ([]datasource.Instructor, error) {
	return nil, nil
}

func (m *mock) GetCoursesByID(id int) (*datasource.Course, error) {
	for _, course := range m.courses {
		if course.ID == id {
			return &course, nil
		}
	}
	return nil, errors.New("bad stuff happened")
}
func (m *mock) GetInstructorByID(id int) (*datasource.Instructor, error) {
	return nil, nil
}
func (m *mock) GetUserByID(id int) (*datasource.User, error) {
	return nil, nil
}

func (m *mock) GetCoursesForInstructor(id int) ([]datasource.Course, error) {
	return nil, nil
}
func (m *mock) GetCoursesForUser(id int) ([]datasource.Course, error) {
	return nil, nil
}

func (m *mock) CreateNewUser(*datasource.User) (int, error) {
	return -1, nil
}
func (m *mock) AddUserInterest(id int, interests []string) (int, error) {
	return -1, nil
}
