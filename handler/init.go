package handler

import (
	"net/http"

	store "paper-test/storage"
)

type IHandler interface {
	Disburse(w http.ResponseWriter, r *http.Request)
	DisburseList(w http.ResponseWriter, r *http.Request)
	UserDetail(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	IHandler
	Storage store.IStorage
}

func NewHandler(storage store.IStorage) *Handler {
	return &Handler{
		Storage: storage,
	}
}
