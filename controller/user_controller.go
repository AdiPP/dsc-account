package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/mock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	helpers.SendResponse(w, r, mock.Users, http.StatusOK)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	vars := mux.Vars(r)

	for _, item := range mock.Users {
		if item.ID == vars["user"] {
			u = item
		}
	}

	if (reflect.DeepEqual(u, entity.User{})) {
		helpers.SendResponse(w, r, u, http.StatusNotFound)
		return
	}

	helpers.SendResponse(w, r, u, http.StatusFound)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		panic(err)
	}

	fails := validateUserRequestBody(u)

	if len(fails) != 0 {
		helpers.SendResponse(w, r, fails, http.StatusBadRequest)
		return
	}

	u.ID = uuid.NewString()
	mock.Users = append(mock.Users, u)

	helpers.SendResponse(w, r, u, http.StatusCreated)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}
	uIdx := int(0)
	vars := mux.Vars(r)

	for i, item := range mock.Users {
		if item.ID == vars["user"] {
			u = item
			uIdx = i
		}
	}

	if (reflect.DeepEqual(u, entity.User{})) {
		helpers.SendResponse(w, r, u, http.StatusNotFound)
		return
	}

	nu := entity.User{}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&nu); err != nil {
		panic(err)
	}

	fails := validateUserRequestBody(nu)

	if len(fails) != 0 {
		helpers.SendResponse(w, r, fails, http.StatusBadRequest)
		return
	}

	u.Username = nu.Username
	u.Password = nu.Password
	u.Email = nu.Email
	u.Name = nu.Name
	u.Roles = nu.Roles

	mock.Users = append(mock.Users, u)
	mock.Users = append(mock.Users[:uIdx], mock.Users[uIdx+1:]...)

	helpers.SendResponse(w, r, u, http.StatusOK)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}
	uIdx := int(0)
	vars := mux.Vars(r)

	for i, item := range mock.Users {
		if item.ID == vars["user"] {
			u = item
			uIdx = i
		}
	}

	if (reflect.DeepEqual(u, entity.User{})) {
		helpers.SendResponse(w, r, u, http.StatusNotFound)
		return
	}

	mock.Users = append(mock.Users[:uIdx], mock.Users[uIdx+1:]...)

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
