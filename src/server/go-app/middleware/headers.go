package middleware

import (
	"net/http"
)

func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		enableCors(limitBodySize(next)).ServeHTTP(res, req)
	})
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		headers := req.Header.Get("Access-Control-Request-Headers")
		(res).Header().Set("Access-Control-Allow-Headers", headers)
		(res).Header().Set("Access-Control-Allow-Credentials", "true")
		(res).Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		if req.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(res, req)
	})
}

func limitBodySize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//check what type of request and set body size

		req.Body = http.MaxBytesReader(res, req.Body, 20000)
		next.ServeHTTP(res, req)
	})
}
