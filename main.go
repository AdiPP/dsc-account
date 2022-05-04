package main

import (
	"github.com/AdiPP/dsc-account/controller"
	router "github.com/AdiPP/dsc-account/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	pingController controller.PingController = controller.NewPingController()
)

func main() {
	const port string = "8080"

	httpRouter.Get("/api/ping", pingController.Ping)

	httpRouter.Serve(port)
}
