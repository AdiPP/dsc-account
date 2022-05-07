package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Roles     []Role         `json:"roles" gorm:"many2many:role_users;constraint:OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	return
}

func (u *User) HasRole(role string) bool {
	for _, val := range u.Roles {
		if string(val.Name) == role {
			return true
		}
	}

	return false
}

func (u *User) HasAnyRoles(roles ...string) bool {
	for _, val := range roles {
		if u.HasRole(val) {
			return true
		}
	}

	return false
}
