package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/repository"
	"github.com/gorilla/mux"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
)

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	usrs := userRepository.FindAll()

	helpers.SendResponse(w, r, usrs, http.StatusOK)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u, err := userRepository.FindOrFail(vars["user"])

	if err != nil {
		helpers.SendResponse(w, r, nil, http.StatusNotFound)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusFound)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}

	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&u)

	fails := validateUserRequestBody(u)

	if len(fails) != 0 {
		helpers.SendResponse(w, r, fails, http.StatusBadRequest)
		return
	}

	u, err := userRepository.Save(u)

	if err != nil {
		helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusCreated)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, err := userRepository.FindOrFail(vars["user"])

	if err != nil {
		helpers.SendResponse(w, r, nil, http.StatusNotFound)
		return
	}

	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&u)

	fails := validateUserRequestBody(u)

	if len(fails) != 0 {
		helpers.SendResponse(w, r, fails, http.StatusBadRequest)
		return
	}

	u, _ = userRepository.Update(u)

	helpers.SendResponse(w, r, u, http.StatusOK)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, err := userRepository.FindOrFail(vars["user"])

	if err != nil {
		helpers.SendResponse(w, r, nil, http.StatusNotFound)
		return
	}

	u, _ = userRepository.Delete(u)

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
