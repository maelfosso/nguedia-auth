package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *HttpServer) Forget(w http.ResponseWriter, r *http.Request) {
	var params ForgetPasswordParams

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
		JSON(w, http.StatusBadRequest, err)
		return
	}

	h.DB.DeleteResetPassword(params.Email)

	token, err := ""
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	data, err := h.DB.SaveResetPasswordToken(params.Email, token)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, true)
	return
}

func (h *HttpServer) Reset(w http.ResponseWriter, r *http.Request) {
	var params VerifyResetPasswordTokenParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	resetToken, err := h.DB.FindResetPassword(params.Email, params.Token)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, resetToken != nil)
}

func (h *HttpServer) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var params ResetPasswordParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	if params.Password != params.ConfirmPassword {
		JSON(w, http.StatusBadRequest, fmt.Errorf("password don't match"))
		return
	}

	rp, err := h.DB.FindResetPassword(params.Email, params.Token)
	if err != nil {
		JSON(w, http.StatusInternalServerError, fmt.Errorf(""))
		return
	}

	if rp == nil {
		JSON(w, http.StatusBadRequest, fmt.Errorf("no data concerning this resetting"))
		return
	}

	// TODO: Check if the token has expired
	// If so reject

	user, err := h.DB.FindUserByEmail(params.Email)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	if err := h.DB.UpdateUserPassword(*user, params.Password); err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusOK, true)
}
