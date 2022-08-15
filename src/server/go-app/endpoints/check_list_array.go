package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/datastore/queries"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
)

func CheckListArray(res http.ResponseWriter, req *http.Request) {
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
	arr, err := extractListArray(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to extract user list")
		return
	}
	if queries.HasListArrayBeenModified(user, arr, ctx) {
		encoder.Encode(true)
		return
	}
	encoder.Encode(false)
}

func extractListArray(req *http.Request) ([]string, error) {
	var arr []string
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&arr)
	if err != nil {
		println(err.Error())
		return arr, err
	}
	return arr, nil

}
