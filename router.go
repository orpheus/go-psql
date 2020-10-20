package main

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/", handler)
	r.HandleFunc("/submission", acceptSubmission).Methods("POST")
	r.HandleFunc("/submission", getSubmission).Methods("GET")
	r.HandleFunc("/user", getUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")

	return r
}
