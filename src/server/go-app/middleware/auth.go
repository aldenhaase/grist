package middleware

import (
	"net/http"
	"server/auth"

	"google.golang.org/appengine/v2"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		if !auth.ValidateSessionCookie(req, ctx) {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(res, req)
	})
}
