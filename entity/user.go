package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Roles     []Role         `json:"roles" gorm:"many2many:role_users;constraint:OnDelete:CASCADE;"`
}

type JsonCreateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Roles    []Role `json:"roles"`
}

type JsonUpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Roles    []Role `json:"roles"`
}

func HashPassword(p string) (string, error) {
	return hashPassword(p)
}

func hashPassword(p string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), 14)

	return string(hashedPassword), err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	u.Password, err = hashPassword(u.Password)

	return err
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password, err = hashPassword(u.Password)

	return err
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

func (req *JsonCreateUserRequest) MapToUser() User {
	u := User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
		Roles:    req.Roles,
	}

	return u
}

func (req *JsonUpdateUserRequest) MapToUser() User {
	u := User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
		Roles:    req.Roles,
	}

	return u
}
