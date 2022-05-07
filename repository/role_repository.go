package repository

import (
	"github.com/AdiPP/dsc-account/entity"
	"gorm.io/gorm/clause"
)

type RoleRepository struct{}

func NewRoleRepository() RoleRepository {
	return RoleRepository{}
}

func (rr *RoleRepository) FindAll() []entity.Role {
	rls := []entity.Role{}

	postgresSqlDatabase.DB.Preload(clause.Associations).Find(&rls)

	return rls
}
