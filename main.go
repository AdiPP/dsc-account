package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AdiPP/dsc-account/controller"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/service"
	"github.com/gorilla/mux"
)

var (
	httpRouter      *mux.Router                = mux.NewRouter()
	pingController  controller.PingController  = controller.NewPingController()
	tokenController controller.TokenController = controller.NewTokenController()
	userController  controller.UserController  = controller.NewUserController()
	tokenService    service.TokenService       = service.NewTokenService()
)

func main() {
	const port string = "8080"

	httpRouter.Use(loggingMiddleware)

	apiRoute := httpRouter.PathPrefix("/api").Subrouter()

	// Ping
	pingRoute := apiRoute.Methods(http.MethodGet).Subrouter()
	pingRoute.HandleFunc("/ping", pingController.Ping).Methods(http.MethodGet)

	// Token
	tknRoute := apiRoute.Methods(http.MethodPost).Subrouter()
	tknRoute.HandleFunc("/tokens", tokenController.IssueToken).Methods(http.MethodPost)
	tknRoute.HandleFunc("/tokens/refresh", tokenController.RefreshToken).Methods(http.MethodPost)

	// User
	usrRoute := apiRoute.Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()
	usrRoute.HandleFunc("/users", userController.GetUsers).Methods(http.MethodGet)
	usrRoute.HandleFunc("/users/{user}", userController.GetUser).Methods(http.MethodGet)
	usrRoute.HandleFunc("/users", userController.CreateUser).Methods(http.MethodPost)
	usrRoute.HandleFunc("/users/{user}", userController.UpdateUser).Methods(http.MethodPatch)
	usrRoute.HandleFunc("/users/{user}", userController.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Mux HTTP server running on port", port)
	http.ListenAndServe(":"+port, httpRouter)
	// httpRouter.Serve(port)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func auth(hf http.HandlerFunc) http.HandlerFunc {
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
