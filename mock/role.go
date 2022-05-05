package mock

import (
	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/valueobject"
)

var AdminRole = entity.Role{
	ID:    "91e6db60-e820-4645-bad9-bcd59813f4c7",
	Name:  valueobject.Admin,
	Label: "Administrator",
}

var UserRole = entity.Role{
	ID:    "ad562c72-d7ef-4406-bef0-04b43d796ac0",
	Name:  valueobject.User,
	Label: "User",
}
