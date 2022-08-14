package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/datastore/queries"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
)

func CreateUserList(res http.ResponseWriter, req *http.Request) {
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
		encoder.Encode("could append to user list")
		return
	}

	ctx := appengine.NewContext(req)
	user := cookie.Username
	if queries.DoesListExist(user, listName, ctx) {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("List by this name already exists")
		return
	}
	key, err := queries.CreateUserList(ctx)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("could not create user list")
		return
	}

	err = queries.AddUserList(key, user, listName, ctx)
	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("could append to user list")
		return
	}
	res.WriteHeader(http.StatusCreated)
}
