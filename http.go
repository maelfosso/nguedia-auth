package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type HttpServer struct {
	db *DBQuery
}

func NewHttpServer(db *DBQuery) http.Handler {
	r := mux.NewRouter()

	h := &HttpServer{
		db: db,
	}

	r.HandleFunc("/signup", h.SignUp)
	r.HandleFunc("/signup", h.SignUp)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	handler := cors.AllowAll().Handler(loggedRouter)

	return handler
}
