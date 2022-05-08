package routes

import (
	"net/http"

	"github.com/AdiPP/dsc-account/controller"
	"github.com/AdiPP/dsc-account/middleware"
	"github.com/AdiPP/dsc-account/valueobject"
	"github.com/gorilla/mux"
)

var (
	pingController  controller.PingController  = controller.NewPingController()
	tokenController controller.TokenController = controller.NewTokenController()
	userController  controller.UserController  = controller.NewUserController()
	roleController  controller.RoleController  = controller.NewRoleController()
)

func apiRoutes(apiRoute *mux.Router) {
	// Ping
	pingRoute := apiRoute.Methods(http.MethodGet).Subrouter()

	pingRoute.HandleFunc("/ping", middleware.Middleware(
		http.HandlerFunc(pingController.Ping),
	).ServeHTTP).Methods(http.MethodGet)

	// Token
	tknRoute := apiRoute.Methods(http.MethodPost, http.MethodGet).Subrouter()

	tknRoute.HandleFunc("/tokens", middleware.Middleware(
		http.HandlerFunc(tokenController.IssueToken),
	).ServeHTTP).Methods(http.MethodPost)

	tknRoute.HandleFunc("/tokens/refresh", middleware.Middleware(
		http.HandlerFunc(tokenController.RefreshToken),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodPost)

	tknRoute.HandleFunc("/auth/me", middleware.Middleware(
		http.HandlerFunc(tokenController.AuthMe),
	).ServeHTTP).Methods(http.MethodGet)

	// Roles
	rlRoute := apiRoute.Methods(http.MethodGet).Subrouter()

	rlRoute.HandleFunc("/roles", middleware.Middleware(
		http.HandlerFunc(roleController.GetRoles),
		middleware.HasRoles(string(valueobject.Admin)),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodGet)

	// User
	usrRoute := apiRoute.Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()

	usrRoute.HandleFunc("/users", middleware.Middleware(
		http.HandlerFunc(userController.GetUsers),
		middleware.HasRoles(string(valueobject.Admin)),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodGet)

	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.GetUser),
		middleware.HasRoles(string(valueobject.Admin), string(valueobject.User)),
		middleware.CanShowUser(),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodGet)

	usrRoute.HandleFunc("/users", middleware.Middleware(
		http.HandlerFunc(userController.CreateUser),
		middleware.HasRoles(string(valueobject.Admin)),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodPost)

	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.UpdateUser),
		middleware.HasRoles(string(valueobject.Admin)),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodPatch)

	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.DeleteUser),
		middleware.HasRoles(string(valueobject.Admin)),
		middleware.AuthMiddleware(),
	).ServeHTTP).Methods(http.MethodDelete)
}
