package endpoints

import (
	"net/http"
	"server/auth"
	"server/extractors"
	"server/lystrTypes"
)

func GetRegistrationCookies(res http.ResponseWriter, req *http.Request) {
	_, ErrNoCookie := req.Cookie(lystrTypes.RCookie_t)

	userIP := extractors.ExtractUserIP(req)
	if userIP == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if ErrNoCookie == nil {
		return
	}
	cookie := auth.GenerateRegCookie(userIP)

	http.SetCookie(res, cookie)
}
