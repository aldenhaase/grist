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

func CheckForUpdates(res http.ResponseWriter, req *http.Request) {
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
	info, err := extractInfo(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("Failed to extract user list")
		return
	}
	if queries.HasRecordBeenModified(user, info.List_Name, info.Last_Modified, ctx) {
		encoder.Encode(true)
		return
	}
	encoder.Encode(false)
}

func extractInfo(req *http.Request) (types.Check_Time, error) {
	var time types.Check_Time
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&time)
	if err != nil {
		return time, err
	} else {
		return time, nil
	}

}
