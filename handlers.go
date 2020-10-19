package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%d] = %q\n", k, v)
	}
}

func acceptSubmission(w http.ResponseWriter, r *http.Request) {

}

func getSubmission(w http.ResponseWriter, r *http.Request) {

}
