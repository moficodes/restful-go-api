package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	// becasue User is a property of server struct
	// we have access to user at all the server instance
	User User
}

type User struct {
	// this tags after our struct lets us change the represantation in json
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(s.User)
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Accept", "application/json")
	var body User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.User = body
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"update": "ok"}`))
}

func main() {
	s := &server{
		User: User{
			Username: "moficodes",
			Email:    "moficodes@gmail.com",
			Age:      27,
		},
	}
	r := mux.NewRouter()
	r.Handle("/", s)
	r.HandleFunc("/user", s.getUser).Methods(http.MethodGet)
	r.HandleFunc("/user", s.updateUser).Methods(http.MethodPut)

	port := "7999"
	log.Println("starting web server on port", port)

	// instead of using the default handler that comes with net/http we use the mux router from gorilla mux
	log.Fatal(http.ListenAndServe(":"+port, r))
}
