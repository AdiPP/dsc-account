package main

import (
	"github.com/AdiPP/dsc-account/controller"
	router "github.com/AdiPP/dsc-account/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	pingController controller.PingController = controller.NewPingController()
	userController controller.UserController = controller.NewUserController()
)

func main() {
	const port string = "8080"

	httpRouter.Get("/api/ping", pingController.Ping)

	// User
	httpRouter.Get("/api/users", userController.GetUsers)
	httpRouter.Get("/api/users/{user}", userController.GetUser)
	httpRouter.Post("/api/users", userController.CreateUser)
	httpRouter.Put("/api/users/{user}", userController.UpdateUser)
	httpRouter.Delete("/api/users/{user}", userController.DeleteUser)

	httpRouter.Serve(port)
}
