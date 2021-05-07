package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lohon.cm/msvc/auth/db"
	"lohon.cm/msvc/auth/utils"
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

func (h *HttpServer) SignIn(w http.ResponseWriter, r *http.Request) {
	var params SignInParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.DB.FindUserByEmail(params.Email)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		JSON(w, http.StatusBadRequest, fmt.Errorf("no user with that email"))
		return
	}

	res := utils.CheckPassword(user.Password, params.Password)
	if !res {
		JSON(w, http.StatusBadRequest, fmt.Errorf("password don't match"))
		return
	}

	token, err := utils.GetJWT(*user)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	response := SignInResponse{
		token,
		*user,
	}

	JSON(w, http.StatusOK, response)
}
