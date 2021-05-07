package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"lohon.cm/msvc/auth/utils"
)

func (h *HttpServer) Forgot(w http.ResponseWriter, r *http.Request) {
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

	token, err := utils.GenerateRandomString(32)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	_, err = h.DB.SaveResetPasswordToken(params.Email, token)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, true)
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

func (h *HttpServer) Verify(w http.ResponseWriter, r *http.Request) {
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

	if time.Now().After(rp.ExpiredAt) {
		JSON(w, http.StatusBadRequest, fmt.Errorf("token has expired"))
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

	if err := h.DB.UpdateUserPassword(*user, params.Password); err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusOK, true)
}
