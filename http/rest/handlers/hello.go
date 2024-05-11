package handlers

import (
	"net/http"
)

// HelloWorld is a handler for hello world
func (s service) HelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, "Hello, World!", http.StatusOK)
	}
}
