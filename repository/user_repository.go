package repository

import (
	"github.com/AdiPP/dsc-account/database"
	"github.com/AdiPP/dsc-account/entity"
	"gorm.io/gorm/clause"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

var (
	postgresSqlDatabase database.PostgresSqlDatabase = database.NewPostgresSqlDatabase()
)

func (ur *UserRepository) Find(id string) entity.User {
	u := entity.User{}

	postgresSqlDatabase.DB.Preload(clause.Associations).First(&u, "id = ?", id)

	return u
}

func (ur *UserRepository) FindByUsername(username string) entity.User {
	u := entity.User{}

	postgresSqlDatabase.DB.Preload(clause.Associations).First(&u, "username = ?", username)

	return u
}

func (ur *UserRepository) FindOrFail(id string) (entity.User, error) {
	u := entity.User{}

	result := postgresSqlDatabase.DB.Preload(clause.Associations).First(&u, "id = ?", id)

	return u, result.Error
}

func (ur *UserRepository) FindByUsernameOrFail(username string) (entity.User, error) {
	u := entity.User{}

	result := postgresSqlDatabase.DB.Preload(clause.Associations).First(&u, "username = ?", username)

	return u, result.Error
}

func (ur *UserRepository) FindAll() []entity.User {
	usrs := []entity.User{}

	postgresSqlDatabase.DB.Preload(clause.Associations).Find(&usrs)

	return usrs
}

func (ur *UserRepository) Save(u entity.User) (entity.User, error) {
	postgresSqlDatabase.DB.Save(&u)

	return u, nil
}

func (ur *UserRepository) Update(u entity.User) (entity.User, error) {
	postgresSqlDatabase.DB.Model(&u).Omit("CreatedAt").Updates(u)

	return u, nil
}

func (ur *UserRepository) Delete(u entity.User) (entity.User, error) {
	postgresSqlDatabase.DB.Delete(&u)

	return u, nil
}
