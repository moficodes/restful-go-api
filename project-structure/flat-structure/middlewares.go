package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// auth token have the structure `bearer <token>`
		// so we split it on the ` ` (space character)
		splitToken := strings.Split(authorization, " ")
		// if we end up with a array of size 2 we have the token as the
		// 2nd item in the array
		if len(splitToken) != 2 {
			// we got something different
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "not authorized"}`))
			return
		}
		// second item is our possible token
		jwtToken := splitToken[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("very-secret"), nil
		})

		if err != nil {
			// we got something different
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "not authorized"}`))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), ContextKey("props"), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "not authorized"}`))
		}

	})
}
