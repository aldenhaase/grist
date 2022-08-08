package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/crypto"
	"server/datastore/queries"
	"server/types"
	"strings"
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
			expiration := time.Now().AddDate(0, 0, 15)
			formatedExpiration := expiration.Format(time.RFC3339)
			signature, err := generateAuthSignature(userInfo.Username, formatedExpiration)
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			val := serializeAuthVals(&types.Authentication_Cookie{
				Username:   userInfo.Username,
				Expiration: formatedExpiration,
				Signature:  signature,
			})
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			cookie := &http.Cookie{
				Name:     "LAUTH",
				Value:    val,
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

func concatinateAuthString(userIP string, expiration string) string {
	return userIP + expiration + os.Getenv("SERVER_SIG")
}

func generateAuthSignature(userIP string, expiration string) (string, error) {
	return crypto.HashPass(concatinateAuthString(userIP, expiration))
}

func serializeAuthVals(source *types.Authentication_Cookie) string {
	return source.Username + "|" + source.Expiration + "|" + source.Signature
}

func deserializeAuthVals(source string) *types.Authentication_Cookie {
	values := strings.Split(source, "|")
	return &types.Authentication_Cookie{
		Username:   values[0],
		Expiration: values[1],
		Signature:  values[2],
	}
}
