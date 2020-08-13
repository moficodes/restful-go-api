package handler

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/moficodes/restful-go-api/pkg/internal/middleware"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(middleware.ContextKey("props")).(jwt.MapClaims)

	name, ok := data["name"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "not authorized"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"message\": \"hello %v\"}", name)))
}
