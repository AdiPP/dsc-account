package controller

import (
	"net/http"

	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/mock"
)

type PingController struct{}

func NewPingController() PingController {
	return PingController{}
}

func (pc *PingController) Ping(w http.ResponseWriter, r *http.Request) {
	helpers.SendResponse(w, r, mock.Ping, http.StatusOK)
}
