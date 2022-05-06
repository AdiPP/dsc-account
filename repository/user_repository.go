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

func (ur *UserRepository) Find(id string) (entity.User, int) {
	for i, item := range mock.Users {
		if item.ID == id {
			return item, i
		}
	}

	return entity.User{}, 0
}

func (ur *UserRepository) FindOrFail(id string) (entity.User, int, error) {
	u, idx := ur.Find(id)

	if !reflect.DeepEqual(u, entity.User{}) {
		return u, 0, nil
	}

	return u, idx, errors.New("user does not exists")
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
	u, idx, err := ur.FindOrFail(u.ID)

	if err != nil {
		return entity.User{}, err
	}

	mock.Users = append(mock.Users[:idx], mock.Users[idx+1:]...)
	mock.Users = append(mock.Users, u)

	return u, nil
}

func (ur *UserRepository) Delete(u entity.User) (entity.User, error) {
	u, idx, err := ur.FindOrFail(u.ID)

	if err != nil {
		return entity.User{}, err
	}

	mock.Users = append(mock.Users[:idx], mock.Users[idx+1:]...)
	return u, nil
}
