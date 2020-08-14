package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func init() {
	loadCourses("../../data/courses.json")
}

func TestGetAllCourses_success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/users", nil)

	GetAllCourses(w, r)
	resp := w.Result()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("readAll err=%s; want nil", err)
	}
	var res []Course
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("unmarshal err=%s; want nil", err)
	}

	want := 100
	got := len(res)

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}
}

func TestGetAllCourses_server(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetAllCourses))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("get error=%s, wanted nil", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("readAll err=%s; want nil", err)
	}
	var res []Course
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("unmarshal err=%s; want nil", err)
	}
	want := 100
	got := len(res)

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}
}

func TestGetCoursesByID_success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/courses/1", nil)

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	GetCoursesByID(w, r)
	resp := w.Result()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("readAll err=%s; want nil", err)
	}
	var res Course
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("unmarshal err=%s; want nil", err)
	}

	want := 1
	got := res.ID

	if want != got {
		t.Errorf("want=1; got=%d", got)
	}
}

func TestGetCoursesByID_failure(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/courses/101", nil)

	vars := map[string]string{
		"id": "101",
	}

	r = mux.SetURLVars(r, vars)

	GetCoursesByID(w, r)
	resp := w.Result()
	want := 404
	got := resp.StatusCode
	if want != got {
		t.Errorf("want=%d; got=%d", want, got)
	}
}

func TestGetCoursesWithInstructorAndAttendee_success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/?instructor=1&attendee=2", nil)

	GetCoursesWithInstructorAndAttendee(w, r)
	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("readAll err=%s; want nil", err)
	}

	var res []Course
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("unmarshal err=%s; want nil", err)
	}
	want := 1
	got := len(res)

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}

	want = 1
	got = res[0].InstructorID

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}
}

func TestGetCoursesWithInstructorAndAttendee_server(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetCoursesWithInstructorAndAttendee))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "?instructor=1&attendee=2")
	if err != nil {
		t.Errorf("get error=%s, wanted nil", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("readAll err=%s; want nil", err)
	}

	var res []Course
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("unmarshal err=%s; want nil", err)
	}
	want := 1
	got := len(res)

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}

	want = 1
	got = res[0].InstructorID

	if err != nil {
		t.Errorf("want=%d; got=%d", want, got)
	}
}
