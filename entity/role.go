package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	Label     string         `json:"label"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}

	return
}
