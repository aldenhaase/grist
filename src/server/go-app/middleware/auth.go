package middleware

import (
	"net/http"
	"server/auth"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !auth.ValidateSessionCookie(req) {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(res, req)
	})
}
