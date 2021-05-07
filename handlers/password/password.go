package password

import "net/http"

func Forget(w http.ResponseWriter, r *http.Request) {}

func Reset(w http.ResponseWriter, r *http.Request) {}

func ValidateToken(w http.ResponseWriter, r *http.Request) {}
