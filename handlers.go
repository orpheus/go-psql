package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/orpheus/wsd/persistence"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%v] = %q\n", k, v)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func createUser(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newUser := &persistence.User{
			ID:        1,
			FirstName: "Ryan",
			LastName:  "Chacon",
			Email:     "ryan.gnar@yahoo.com",
		}
		persistence.SaveUser(db, newUser)
	}
}
