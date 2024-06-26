package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware)
	// hello world

	r.Use(handler.MiddlewareLogger())
	r.HandleFunc("/", handler.HelloWorld()).Methods(http.MethodGet)
	r.HandleFunc("/healthz", handler.Health())
	r.HandleFunc("/todo", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", handler.Delete()).Methods(http.MethodDelete)
}
