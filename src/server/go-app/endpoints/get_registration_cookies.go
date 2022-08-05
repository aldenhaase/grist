package endpoints

import (
	"net/http"
)

func GetRegistrationCookies(res http.ResponseWriter, req *http.Request) {
	_, cookieExists := checkForRegistrationCookie(res, req)

	var userIP = extractUserIP(req)
	if userIP == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if cookieExists {
		return
	}
	cookie, err := generateCookie(userIP)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(res, cookie)
}
