package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/errors"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/repository"
	"github.com/AdiPP/dsc-account/service"
	"github.com/gorilla/mux"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUserService()
)

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	usrs := userRepository.FindAll()

	helpers.SendResponse(w, r, usrs, http.StatusOK)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u, err := userRepository.FindOrFail(vars["user"])

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusNotFound)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusFound)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := entity.JsonCreateUserRequest{}

	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&req)

	u := req.MapToUser()

	fails := validateUserRequestBody(u)

	if len(fails) != 0 {
		es := errors.NewRequestValidationError(fails, http.StatusBadRequest)
		helpers.SendResponse(w, r, es.Message, es.StatusCode)
		return
	}

	u, err := userService.Create(u)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusCreated)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := entity.JsonUpdateUserRequest{}

	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&req)

	req.ID = vars["user"]

	_, err := userRepository.FindOrFail(req.ID)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusNotFound)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	u := req.MapToUser()

	fails := validateUserRequestBody(u)

	if len(fails) != 0 {
		es := errors.NewRequestValidationError(fails, http.StatusBadRequest)
		helpers.SendResponse(w, r, es.Message, es.StatusCode)
		return
	}

	u, err = userService.Update(u)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusOK)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, err := userRepository.FindOrFail(vars["user"])

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusNotFound)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	u, err = userService.Delete(u)

	if err != nil {
		es := errors.NewServiceError(err.Error(), http.StatusInternalServerError)
		helpers.SendResponse(w, r, es, es.StatusCode)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusOK)
}

func validateUserRequestBody(u entity.User) url.Values {
	fails := url.Values{}

	if u.Username == "" {
		fails.Add("username", "The username field is required!")
	}

	if u.Password == "" {
		fails.Add("password", "The password field is required!")
	}

	if u.Email == "" {
		fails.Add("email", "The email field is required!")
	}

	if u.Name == "" {
		fails.Add("name", "The name field is required!")
	}

	return fails
}
