package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/helpers"
	"github.com/AdiPP/dsc-account/mock"
	"github.com/AdiPP/dsc-account/service"
	"github.com/golang-jwt/jwt/v4"
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

func AuthMiddleware() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtTknStr, err := getToken(r)

			if err != nil {
				helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
				return
			}

			_, err = tokenService.ValidateToken(jwtTknStr)

			if err != nil {
				helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func HasRoles(roles ...string) MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtTknStr, _ := getToken(r)

			jwtTkn, _ := jwt.Parse(jwtTknStr, func(token *jwt.Token) (interface{}, error) {
				return service.JwtKey, nil
			})

			claims, ok := jwtTkn.Claims.(jwt.MapClaims)

			if !ok || !jwtTkn.Valid {
				helpers.SendResponse(w, r, nil, http.StatusUnauthorized)
				return
			}

			usrnm := claims["username"]
			u := entity.User{}

			for _, item := range mock.Users {
				if item.Username == usrnm {
					u = item
				}
			}

			if !u.HasAnyRoles(roles...) {
				helpers.SendResponse(w, r, nil, http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getToken(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		return "", errors.New("unauthorized")
	}

	token := strings.TrimSpace(splitToken[1])

	return token, nil
}
