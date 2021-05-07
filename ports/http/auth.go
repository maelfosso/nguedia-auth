package http

import (
	"encoding/json"
	"net/http"

	"lohon.cm/msvc/auth/db"
)

func (h *HttpServer) SignUp(w http.ResponseWriter, r *http.Request) {
	var user *db.User

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	foundUser, err := h.DB.FindUserByEmail(user.Email)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	if foundUser != nil {
		JSON(w, http.StatusBadRequest, "user with this email already exists")
		return
	}

	if err := h.DB.SaveUser(user); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Send message through the broker annoucing the user creation
	// TODO: Send email - User account Verification

	JSON(w, http.StatusOK, true)
}

func (hs *HttpServer) SignIn(w http.ResponseWriter, r *http.Request) {

}
