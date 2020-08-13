package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	users []User
)

func init() {
	loadUsers("./data/users.json")
}

func loadUsers(path string) {
	if err := readContent(path, &users); err != nil {
		log.Println("Could not read instructors data")
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
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
