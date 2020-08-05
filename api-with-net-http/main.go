package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// A REST server has 3 main components

// Routes
// Route represent the State in REST
// As we know now, REST deals with the represantation of data
// in simplest term, route is anything after the host
// we match these routes in different ways to handle http actions

// Handlers
// Handlers as the name suggests, handles.
// In go terms Handler is an interface that has only one method: ServeHTTP
// It is any function that takes a ResponseWriter and pointer to a Request

// Server
// Server is the workhorse for our application
// this is what takes care of incoming request from our client
// It gets many names Mux, ServerMux, Server, Router etc.

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

//
func (s *server) user(w http.ResponseWriter, r *http.Request) {
	// just because we are writing JSON does not mean our client will understand
	// with this header we make it explicit
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	// if request method is GET business as usual
	case "GET":
		e := json.NewEncoder(w)
		e.Encode(s.User)
	case "PUT":
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
	// for all other query
	// return empty response and 404 status code
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// this handlerFuncs job is simple
// send back whatever was sent to it as base64 encoded
func getBase64(w http.ResponseWriter, r *http.Request) {
	message := strings.Split(r.URL.String(), "/")[2]
	data := []byte(message)
	str := base64.StdEncoding.EncodeToString(data)
	w.Write([]byte(str))
}

func main() {
	s := &server{
		User: User{
			Username: "moficodes",
			Email:    "moficodes@gmail.com",
			Age:      27,
		},
	}
	// because s is an instance on server it is now a handler and we can pass it to http.Handle
	http.Handle("/", s)
	http.HandleFunc("/user", s.user)
	http.HandleFunc("/base64/", getBase64)

	port := "7999"
	log.Println("starting web server on port", port)
	// this is a blocking process
	// go will wait for requests to come and program will not exit
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
