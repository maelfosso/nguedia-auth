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

type ForgetPasswordParams struct {
	Email string `json:"email"`
}

type VerifyResetPasswordTokenParams struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResetPasswordParams struct {
	Email           string `json:"email"`
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
