package valueobject

type Role string

const (
	Unknown Role = ""
	Admin   Role = "admin"
	User    Role = "user"
)
