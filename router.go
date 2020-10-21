package main

import (
	"database/sql"
	"github.com/gorilla/mux"
)

func createRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/", handler)
	r.HandleFunc("/user", getUser).Methods("GET")
	r.HandleFunc("/user", createUser(db)).Methods("POST")

	return r
}
