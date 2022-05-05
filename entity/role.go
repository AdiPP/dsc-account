package entity

import "github.com/AdiPP/dsc-account/valueobject"

type Role struct {
	ID    string           `json:"id"`
	Name  valueobject.Role `json:"name"`
	Label string           `json:"label"`
}
