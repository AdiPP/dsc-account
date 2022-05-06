package service

import (
	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/repository"
)

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
)

func (us *UserService) Create(u entity.User) (entity.User, error) {
	u, err := userRepository.Save(u)

	return u, err
}

func (us *UserService) Update(u entity.User) (entity.User, error) {
	u, err := userRepository.Update(u)

	return u, err
}

func (us *UserService) Delete(u entity.User) (entity.User, error) {
	u, err := userRepository.Delete(u)

	return u, err
}
