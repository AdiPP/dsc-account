package main

import (
	"github.com/AdiPP/dsc-account/database"
	"github.com/AdiPP/dsc-account/middleware"
	"github.com/AdiPP/dsc-account/routes"
)

var (
	muxRouter routes.MuxRouter  = routes.NewMuxRouter()
	db        database.Database = database.NewDatabase()
)

func init() {
	db.Init()
	muxRouter.InitMiddleware(middleware.LoggingMiddleware)
	muxRouter.InitApiRoutes()
}

func main() {
	const port string = "8080"

	muxRouter.Serve(port)
}
