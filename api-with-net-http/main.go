package main

import (
	"log"
	"net/http"
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

type anything int

// to be considered a handler, a type needs to implement the ServeHTTP method
func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}

func main() {
	var thisthing anything
	// because thisthing is an instance on anything it is now a handler and we can pass it to http.Handle
	http.Handle("/", thisthing)
	port := "7999"
	log.Println("starting web server on port", port)
	// this is a blocking process
	// go will wait for requests to come and program will not exit
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
