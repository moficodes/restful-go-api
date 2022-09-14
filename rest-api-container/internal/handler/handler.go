package handler

import (
	"github.com/moficodes/restful-go-api/database/internal/datasource"
)

func NewHandler(db datasource.DB) *Handler {
	h := Handler{db}
	return &h
}
