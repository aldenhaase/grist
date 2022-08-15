package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/datastore/queries"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
)

func DeleteUserList(res http.ResponseWriter, req *http.Request) {
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

	var listName string
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&listName)

	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("could not extract list name")
		return
	}

	ctx := appengine.NewContext(req)
	user := cookie.Username

	if !queries.DoesListExist(user, listName, ctx) {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("List does not exitst")
		return
	}

	err = queries.DeleteUserList(user, listName, ctx)
	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("could not delete user list")
		return
	}
}
