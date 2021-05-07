package http

import (
	"encoding/json"
	"net/http"

	"lohon.cm/msvc/auth/db"
)

type HttpServer struct {
	DB *db.DBQuery
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
