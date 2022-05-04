package mock

import (
	"github.com/AdiPP/dsc-account/entity"
)

var User = entity.User{
	ID:       "4f9d7872-85ff-40e6-b068-232ab1b009da",
	Username: "dummy",
	Password: "dummy1234",
	Email:    "dummy@gmail.com",
	Name:     "Dummy",
}

var Users = []entity.User{
	User,
}
