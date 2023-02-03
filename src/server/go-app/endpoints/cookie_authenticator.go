package endpoints

import (
	"encoding/json"
	"net/http"
	"server/auth"

	"google.golang.org/appengine/v2"
)

func AuthenticateCookie(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	encoder := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")
	if !auth.ValidateSessionCookie(req, ctx) {
		res.WriteHeader(http.StatusOK)
		encoder.Encode(false)
		return
	}
	res.WriteHeader(http.StatusOK)
	encoder.Encode(true)
}
