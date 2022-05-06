package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/service"
)

type MiddlewareAdapter func(http.Handler) http.Handler

var (
	tokenService service.TokenService = service.NewTokenService()
)

func Middleware(handler http.Handler, middlewareAdapters ...MiddlewareAdapter) http.Handler {
	for i := len(middlewareAdapters); i > 0; i-- {
		handler = middlewareAdapters[i-1](handler)
	}

	return handler
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Auth() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			next.ServeHTTP(w, r)
		})
	}
}
