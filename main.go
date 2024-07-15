package main

import (
	"fmt"
	"net/http"

	hand "paper-test/handler"
	store "paper-test/storage"

	"github.com/gorilla/mux"
)

func main() {
	// init storage
	storage := store.NewStorage()
	storage.InitData()

	// init handler
	handler := &hand.Handler{
		Storage: storage,
	}

	// init router
	router := mux.NewRouter()

	// main router
	router.HandleFunc("/disburse", handler.Auth(handler.Disburse)).Methods("POST")
	router.HandleFunc("/disburse_list", handler.Auth(handler.DisburseList)).Methods("GET")
	router.HandleFunc("/user_detail", handler.Auth(handler.UserDetail)).Methods("GET")

	fmt.Println("listening server on port :8080")
	http.ListenAndServe(":8080", router)
}
