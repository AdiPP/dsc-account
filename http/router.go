package http

import "net/http"

type Router interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Put(uri string, f func(w http.ResponseWriter, r *http.Request))
	Patch(uri string, f func(w http.ResponseWriter, r *http.Request))
	Delete(uri string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}