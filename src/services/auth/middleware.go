package auth

import (
	"net/http"
	"strings"
)

type MiddlewareAuthOptions struct {
	AdminRequired bool
	UserRequired  bool
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	opt := MiddlewareAuthOptions{AdminRequired: true, UserRequired: true}
	return authenticationMiddleWare(next, opt)
}

func RequireUser(next http.HandlerFunc) http.HandlerFunc {
	opt := MiddlewareAuthOptions{AdminRequired: false, UserRequired: true}
	return authenticationMiddleWare(next, opt)
}

func authenticationMiddleWare(next http.HandlerFunc, opt MiddlewareAuthOptions) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !opt.AdminRequired && !opt.UserRequired {
			next(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if isValidToken(token) {
			next(w, r)
			return
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}

func isValidToken(token string) bool {
	_ = token
	return false
}
