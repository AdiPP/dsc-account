package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct{}

var (
	dispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &MuxRouter{}
}

func (m *MuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	dispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	dispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (m *MuxRouter) Put(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	dispatcher.HandleFunc(uri, f).Methods(http.MethodPut)
}

func (m *MuxRouter) Patch(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	dispatcher.HandleFunc(uri, f).Methods(http.MethodPatch)
}

func (m *MuxRouter) Delete(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	dispatcher.HandleFunc(uri, f).Methods(http.MethodDelete)
}

func (m *MuxRouter) Serve(port string) {
	fmt.Println("Mux HTTP server running on port ", port)
	http.ListenAndServe(":"+port, dispatcher)
}
