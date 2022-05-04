package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/dsc-account/mock"
)

type PingController struct{}

func NewPingController() PingController {
	return PingController{}
}

func (pc *PingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mock.Ping)
}
