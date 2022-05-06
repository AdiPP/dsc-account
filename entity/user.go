package entity

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Roles    []Role `json:"roles"`
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
