package http

import "net/http"

type Router interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}
