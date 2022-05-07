package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AdiPP/dsc-account/errors"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/service"
)

type TokenController struct{}

func NewTokenController() TokenController {
	return TokenController{}
}

var (
	tokenService service.TokenService = service.NewTokenService()
)

func (tc *TokenController) AuthMe(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		es := errors.NewServiceError("Unauthorized", http.StatusUnauthorized)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	jwtTknStr := strings.TrimSpace(splitToken[1])

	u, err := tokenService.AuthUser(jwtTknStr)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusOK)
}

func (tc *TokenController) IssueToken(w http.ResponseWriter, r *http.Request) {
	crdn := service.Credential{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&crdn)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusBadRequest)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	tkn, err := tokenService.IssueToken(crdn)

	if err != nil {
		if err.Error() == "credential is invalid" {
			es := errors.NewServiceError(err.Error(), http.StatusUnauthorized)
			helpers.SendResponse(w, r, es, es.StatusCode)
			return
		}

		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, tkn, http.StatusOK)
}

func (tc *TokenController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		es := errors.NewServiceError("Unauthorized", http.StatusUnauthorized)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	jwtTknStr := strings.TrimSpace(splitToken[1])

	tkn, err := tokenService.RefreshToken(jwtTknStr)

	if err != nil {
		if err.Error() == "bad request" {
			es := errors.NewServiceError(err.Error(), http.StatusBadRequest)
			helpers.SendResponse(w, r, es, es.StatusCode)
			return
		}

		if err.Error() == "internal server error" {
			es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
			helpers.SendResponse(w, r, es, es.StatusCode)
			return
		}

		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, tkn, http.StatusOK)
}
