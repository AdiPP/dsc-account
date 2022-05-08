package mock

import (
	"github.com/AdiPP/dsc-account/entity"
)

var AdminUser = entity.User{
	ID:       "4f9d7872-85ff-40e6-b068-232ab1b009da",
	Username: "admin",
	Password: "avada_kedavra",
	Email:    "admin@mail.com",
	Name:     "Administrator",
	Roles:    []entity.Role{AdminRole},
}

var BasicUser = entity.User{
	ID:       "88fe2abc-28e9-4200-91a7-520696a456cf",
	Username: "basic_user",
	Password: "capacious_extremis",
	Email:    "basic_user@mail.com",
	Name:     "Basic User",
	Roles:    []entity.Role{UserRole},
}

var Users = []entity.User{
	AdminUser,
	BasicUser,
}
