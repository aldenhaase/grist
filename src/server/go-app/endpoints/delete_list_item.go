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

func DeleteListItem(res http.ResponseWriter, req *http.Request) {
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
	requestList, err := extractDeleteList(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to extract user list")
		return
	}
	list, err := queries.DeleteListItem(user, ctx, requestList)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to delete item")
		return
	}
	encoder.Encode(list)
}

func extractDeleteList(req *http.Request) ([]string, error) {
	var listItem types.Delete_List
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&listItem)
	if err != nil {
		return listItem.Items, err
	} else {
		return listItem.Items, nil
	}

}
