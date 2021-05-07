package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/go-hclog"
	"lohon.cm/msvc/auth/db"
)

func init() {
	// Checking if environment variables are available
}

var (
	log = hclog.Default()
)

func main() {
	database := db.Database()
	defer db.CloseDB()
	dbQuery := db.NewQuery(database)

	errChan := make(chan error)

	// HTTP
	go func() {
		handler := NewHttpServer(dbQuery)

		log.Info("Listen to port :3000")
		errChan <- http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler)
	}()

	// gRPC

	// NATS

	log.Error("", <-errChan)
}
