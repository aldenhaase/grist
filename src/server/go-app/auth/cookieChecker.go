package auth

import (
	"net/http"
	"server/deserializers"
	"server/lystrTypes"
	"time"
)

func ValidateSessionCookie(req *http.Request) bool {
	cookie, err := req.Cookie(lystrTypes.SCookie_t)
	if err != nil {
		return false
	}
	sessionC := deserializers.SessionCookie(cookie.Value)
	if !VerifySignature(sessionC.Signature, sessionC.Username, sessionC.Expiration) {
		return false
	}
	return true
}

func ValidateRegistrationCookie(cookie *http.Cookie, userIP string) bool {
	sessionC := deserializers.SessionCookie(cookie.Value)
	expiration, err := time.Parse(time.RFC3339, sessionC.Expiration)
	if err != nil {
		println(err.Error())
		return false
	}
	if time.Now().After(expiration) {
		return false
	}
	return VerifySignature(sessionC.Signature, sessionC.Username, sessionC.Expiration)
}
