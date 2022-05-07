package main

import (
	"github.com/AdiPP/dsc-account/database"
	"github.com/AdiPP/dsc-account/middleware"
	"github.com/AdiPP/dsc-account/routes"
)

var (
	muxRouter           routes.MuxRouter             = routes.NewMuxRouter()
	postgresSqlDatabase database.PostgresSqlDatabase = database.NewPostgresSqlDatabase()
)

func init() {
	postgresSqlDatabase.Init()
	muxRouter.InitMiddleware(middleware.LoggingMiddleware)
	muxRouter.InitApiRoutes()
}

func main() {
	const port string = "8080"

	muxRouter.Serve(port)
}
