package endpoints

import (
	"encoding/json"
	"net/http"
	"server/auth"
	"server/dbFuncs"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func LogIn(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	var userInfo lystrTypes.UserQuery
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&userInfo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dbFuncs.DoesPasswordMatch(ctx, userInfo)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie := auth.GenerateSessionCookie(userInfo)
	http.SetCookie(res, cookie)
}
