package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	Router *mux.Router
}

func NewMuxRouter() MuxRouter {
	mr := MuxRouter{
		Router: mux.NewRouter(),
	}

	return mr
}

func (mr *MuxRouter) InitMiddleware(mws ...mux.MiddlewareFunc) {
	for _, mw := range mws {
		mr.Router.Use(mw)
	}
}

func (mr *MuxRouter) InitApiRoutes() {
	apiRtr := mr.Router.PathPrefix("/api").Subrouter()
	apiRoutes(apiRtr)
}

func (mr *MuxRouter) Serve(port string) {
	fmt.Println("Mux HTTP server running on port", port)
	http.ListenAndServe(":"+port, mr.Router)
}
