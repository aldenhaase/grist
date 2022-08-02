package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/crypto"
	"server/datastore/queries"
	"server/types"
	"time"

	"google.golang.org/appengine/v2"
)

func LogIn(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	var userInfo types.UserRecord
	decoder := json.NewDecoder(req.Body)
	encoder := json.NewEncoder(res)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&userInfo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("failed to decode into userInfo")
		return
	} else {
		err = queries.DoesPasswordMatch(ctx, userInfo)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder.Encode("PassworMatchingFailed")
			encoder.Encode(err.Error())
			return
		} else {
			signature, err := crypto.HashPass(userInfo.Username + "secret Key")
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				encoder.Encode(err)
				return

			}
			log.Println(os.Getenv("DOMAIN"))
			cookie := &http.Cookie{
				Name:     userInfo.Username,
				Value:    signature,
				Expires:  time.Now().AddDate(1, 0, 1),
				Secure:   true,
				HttpOnly: true,
				Domain:   os.Getenv("DOMAIN"),
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(res, cookie)
			encoder.Encode("Authorized")
		}
	}
}
