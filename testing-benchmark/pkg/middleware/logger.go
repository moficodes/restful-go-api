package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func Chain(f http.Handler,
	middlewares ...func(next http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		next.ServeHTTP(w, r)
	})
}

func MuxLogger(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Println(r.URL.Path, time.Since(start)) }()
		next.ServeHTTP(w, r)
	})
}
