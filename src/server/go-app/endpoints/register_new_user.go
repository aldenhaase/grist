package endpoints

import (
	"encoding/json"
	"net/http"
	"server/auth"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func RegisterNewUser(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cookie, ErrNoCookie := req.Cookie(lystrTypes.RCookie_t)
	encoder := json.NewEncoder(res)

	var userIP = extractors.ExtractUserIP(req)
	if userIP == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if ErrNoCookie != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if !auth.ValidateRegistrationCookie(cookie, userIP) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if dbFuncs.HasIpMetQuota(ctx, userIP) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	dbFuncs.AddNewUserToDatabase(res, req, ctx)
	encoder.Encode(true)
}
