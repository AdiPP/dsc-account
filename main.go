package main

import (
	"fmt"
	"net/http"

	"github.com/AdiPP/dsc-account/controller"
	"github.com/AdiPP/dsc-account/middleware"
	"github.com/gorilla/mux"
)

var (
	httpRouter      *mux.Router                = mux.NewRouter()
	pingController  controller.PingController  = controller.NewPingController()
	tokenController controller.TokenController = controller.NewTokenController()
	userController  controller.UserController  = controller.NewUserController()
)

func init() {
	httpRouter.Use(middleware.LoggingMiddleware)
}

func main() {
	const port string = "8080"

	// Api Subrouter
	apiRoute := httpRouter.PathPrefix("/api").Subrouter()

	// Ping
	pingRoute := apiRoute.Methods(http.MethodGet).Subrouter()
	pingRoute.HandleFunc("/ping", middleware.Middleware(
		http.HandlerFunc(pingController.Ping),
	).ServeHTTP).Methods(http.MethodGet)

	// Token
	tknRoute := apiRoute.Methods(http.MethodPost).Subrouter()
	tknRoute.HandleFunc("/tokens", middleware.Middleware(
		http.HandlerFunc(tokenController.IssueToken),
	).ServeHTTP).Methods(http.MethodPost)
	tknRoute.HandleFunc("/tokens/refresh", middleware.Middleware(
		http.HandlerFunc(tokenController.RefreshToken),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodPost)

	// User
	usrRoute := apiRoute.Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()
	usrRoute.HandleFunc("/users", middleware.Middleware(
		http.HandlerFunc(userController.GetUsers),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodGet)
	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.GetUser),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodGet)
	usrRoute.HandleFunc("/users", middleware.Middleware(
		http.HandlerFunc(userController.CreateUser),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodPost)
	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.UpdateUser),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodPatch)
	usrRoute.HandleFunc("/users/{user}", middleware.Middleware(
		http.HandlerFunc(userController.DeleteUser),
		middleware.Auth(),
	).ServeHTTP).Methods(http.MethodDelete)

	fmt.Println("Mux HTTP server running on port", port)
	http.ListenAndServe(":"+port, httpRouter)
}
