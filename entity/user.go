package entity

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Roles    []Role `json:"roles"`
}
