package controller

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (tc *TokenController) IssueToken(w http.ResponseWriter, r *http.Request) {
	crdn := service.Credential{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&crdn)

	if err != nil {
		helpers.SendResponse(w, r, err, http.StatusBadRequest)
		return
	}

	u, err := userRepository.FindByUsernameOrFail(crdn.Username)

	if err != nil {
		helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
		return
	}

	tkn, err := tokenService.IssueToken(u, crdn)

	if err != nil {
		if err.Error() == "credential is invalid" {
			helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
			return
		}

		helpers.SendResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, r, tkn, http.StatusOK)
}

func (tc *TokenController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
		return
	}

	jwtTknStr := strings.TrimSpace(splitToken[1])

	tkn, err := tokenService.RefreshToken(jwtTknStr)

	if err != nil {
		if err.Error() == "bad request" {
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		if err.Error() == "internal server error" {
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
		return
	}

	helpers.SendResponse(w, r, tkn, http.StatusOK)
}
