package handler

// func TestGetAllCourses_success(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodGet, "/courses", nil)
// 	e := echo.New()
// 	c := e.NewContext(r, w)
// 	GetAllCourses(c)
// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Errorf("readAll err=%s; want nil", err)
// 	}
// 	var res []datasource.Course
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		t.Errorf("unmarshal err=%s; want nil", err)
// 	}

// 	want := 100
// 	got := len(res)

// 	if err != nil {
// 		t.Errorf("want=%d; got=%d", want, got)
// 	}
// }

// func TestGetCoursesByID_success(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodGet, "/", nil)
// 	e := echo.New()

// 	c := e.NewContext(r, w)
// 	c.SetPath("/courses/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues("1")

// 	GetCoursesByID(c)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Errorf("readAll err=%s; want nil", err)
// 	}
// 	var res datasource.Course
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		t.Errorf("unmarshal err=%s; want nil", err)
// 	}

// 	want := 1
// 	got := res.ID

// 	if want != got {
// 		t.Errorf("want=1; got=%d", got)
// 	}

// 	want = 201
// 	got = w.Code

// 	if want != got {
// 		t.Errorf("want=%d; got=%d", want, got)
// 	}
// }

// func TestGetCoursesByID_failure(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodGet, "/", nil)

// 	e := echo.New()

// 	c := e.NewContext(r, w)
// 	c.SetPath("/courses/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues("101")
// 	err := GetCoursesByID(c)
// 	he, ok := err.(*echo.HTTPError)
// 	if !ok {
// 		t.Error("should be http error")
// 	}

// 	want := 404
// 	got := he.Code
// 	if want != got {
// 		t.Errorf("want=%d; got=%d", want, got)
// 	}
// }
