package controller

import (
	"net/http"

	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/repository"
)

type RoleController struct{}

func NewRoleController() RoleController {
	return RoleController{}
}

var (
	roleRepository repository.RoleRepository = repository.NewRoleRepository()
)

func (rc *RoleController) GetRoles(w http.ResponseWriter, r *http.Request) {
	rls := roleRepository.FindAll()

	helpers.SendResponse(w, r, rls, http.StatusOK)
}
