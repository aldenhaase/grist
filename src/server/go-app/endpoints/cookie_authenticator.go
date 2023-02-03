package endpoints

import (
	"encoding/json"
	"net/http"
	"server/auth"
)

func AuthenticateCookie(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")
	if !auth.ValidateSessionCookie(req) {
		res.WriteHeader(http.StatusOK)
		encoder.Encode(false)
		return
	}
	res.WriteHeader(http.StatusOK)
	encoder.Encode(true)
}
