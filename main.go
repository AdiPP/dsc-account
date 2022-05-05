package main

import (
	"net/http"
	"strings"

	"github.com/AdiPP/dsc-account/controller"
	"github.com/AdiPP/dsc-account/helpers"
	router "github.com/AdiPP/dsc-account/http"
	"github.com/AdiPP/dsc-account/service"
)

var (
	httpRouter      router.Router              = router.NewMuxRouter()
	pingController  controller.PingController  = controller.NewPingController()
	tokenController controller.TokenController = controller.NewTokenController()
	userController  controller.UserController  = controller.NewUserController()
	tokenService    service.TokenService       = service.NewTokenService()
)

func main() {
	const port string = "8080"

	httpRouter.Get("/api/ping", pingController.Ping)

	// Token
	httpRouter.Post("/api/tokens", tokenController.IssueToken)
	httpRouter.Post("/api/tokens/refresh", Auth(tokenController.RefreshToken))

	// User
	httpRouter.Get("/api/users", Auth(userController.GetUsers))
	httpRouter.Get("/api/users/{user}", Auth(userController.GetUser))
	httpRouter.Post("/api/users", Auth(userController.CreateUser))
	httpRouter.Put("/api/users/{user}", Auth(userController.UpdateUser))
	httpRouter.Delete("/api/users/{user}", Auth(userController.DeleteUser))

	httpRouter.Serve(port)
}

func Auth(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")

		if len(splitToken) != 2 {
			helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
			return
		}

		jwtTknStr := strings.TrimSpace(splitToken[1])

		_, err := tokenService.ValidateToken(jwtTknStr)

		if err != nil {
			helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
			return
		}

		hf.ServeHTTP(w, r)
	}
}
