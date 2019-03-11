package main

import (
	"fmt"
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleAbout() http.HandlerFunc {
	thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About")
	}
}

func prepareThing() string {
	return "coo"
}
