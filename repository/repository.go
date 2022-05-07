package repository

import "github.com/AdiPP/dsc-account/database"

var (
	postgresSqlDatabase database.PostgresSqlDatabase = database.NewPostgresSqlDatabase()
)
