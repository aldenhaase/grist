package endpoints

import (
	"encoding/json"
	"net/http"
	"server/session"
)

func AuthenticateCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("LAUTH")
	encoder := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		res.WriteHeader(http.StatusOK)
		encoder.Encode(session.SessionAuthenticationResponse{Authenticated: false})
		return
	}
	sessionCookie := session.Unpack(cookie.Value)
	if !session.VerifySignature(sessionCookie) {
		res.WriteHeader(http.StatusOK)
		encoder.Encode(session.SessionAuthenticationResponse{Authenticated: false})
		return
	}
	res.WriteHeader(http.StatusOK)
	encoder.Encode(session.SessionAuthenticationResponse{Authenticated: true})
}
