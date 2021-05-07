package http

import "lohon.cm/msvc/auth/db"

type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string  `json:"token"`
	User  db.User `json:"user"`
}
