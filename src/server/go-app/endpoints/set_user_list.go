package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/datastore/queries"
	"server/types"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
)

func SetUserList(res http.ResponseWriter, req *http.Request) {
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
	requestList, err := extractUserList(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to extract user list")
		return
	}
	err = queries.SetUserList(user, ctx, requestList)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to set new item")
		return
	}
}

func extractUserList(req *http.Request) (types.User_List, error) {
	var list types.User_List
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&list)
	if err != nil {
		return list, err
	} else {
		return list, nil
	}

}
