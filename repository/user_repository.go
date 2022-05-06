package repository

import (
	"errors"
	"reflect"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/mock"
	"github.com/google/uuid"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur *UserRepository) Find(id string) entity.User {
	for _, item := range mock.Users {
		if item.ID == id {
			return item
		}
	}

	return entity.User{}
}

func (ur *UserRepository) FindOrFail(id string) (entity.User, error) {
	u := ur.Find(id)

	if !reflect.DeepEqual(u, entity.User{}) {
		return u, nil
	}

	return u, errors.New("user does not exists")
}

func (ur *UserRepository) FindAll() []entity.User {
	return mock.Users
}

func (ur *UserRepository) Save(u entity.User) (entity.User, error) {
	u.ID = uuid.NewString()
	mock.Users = append(mock.Users, u)

	return u, nil
}

func (ur *UserRepository) Update(u entity.User) (entity.User, error) {
	for i, item := range mock.Users {
		if item.ID == u.ID {
			mock.Users = append(mock.Users[:i], mock.Users[i+1:]...)
			mock.Users = append(mock.Users, u)
			return u, nil
		}
	}

	return u, errors.New("user does not exists")
}

func (ur *UserRepository) Delete(u entity.User) (entity.User, error) {
	for i, item := range mock.Users {
		if item.ID == u.ID {
			mock.Users = append(mock.Users[:i], mock.Users[i+1:]...)
			return u, nil
		}
	}

	return u, errors.New("user does not exists")
}
