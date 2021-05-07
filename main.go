package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/rs/cors"
)

func init() {
	// Checking if environment variables are available
}

var (
	log = hclog.Default()
)

func main() {

	r := mux.NewRouter()

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	handler := cors.AllowAll().Handler(loggedRouter)

	log.Info("Listen to port :3000")
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler)
}
