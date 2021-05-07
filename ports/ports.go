package ports

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"lohon.cm/msvc/auth/db"
	phttp "lohon.cm/msvc/auth/ports/http"
)

func NewHttpServer(db *db.DBQuery) http.Handler {
	r := mux.NewRouter()

	h := &phttp.HttpServer{
		DB: db,
	}

	r.HandleFunc("/signup", h.SignUp)
	r.HandleFunc("/signup", h.SignIn)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	handler := cors.AllowAll().Handler(loggedRouter)

	return handler
}

func NewGrpcServer() {

}
