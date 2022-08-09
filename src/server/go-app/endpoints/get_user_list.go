package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/datastore/queries"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
)

func GetUserList(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	cookieArr, err := req.Cookie("LAUTH")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		encoder.Encode("Ya got no cookies")
		return
	}
	cookie := deserializeAuthVals(cookieArr.Value)
	err = bcrypt.CompareHashAndPassword([]byte(cookie.Signature), []byte(cookie.Username+cookie.Expiration+os.Getenv("SERVER_SIG")))
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		encoder.Encode("Password does not match")
		return
	}
	ctx := appengine.NewContext(req)
	user := cookie.Username
	list, err := queries.GetUserList(user, ctx)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Could Not Query Userlist")
		return
	}
	encoder.Encode(list)
}
