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

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/datastore"
)

const RegCookieName = "LREG"

func RegisterNewUser(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cookie, cookieExists := checkForRegistrationCookie(res, req)
	var userIP = extractUserIP(req)
	if userIP == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if !cookieExists {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if !validateRegistrationCookie(cookie, userIP) {
		res.WriteHeader(http.StatusConflict)
		return
	}
	if queries.HasIpMetQuota(ctx, userIP) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	addNewUserToDatabase(res, req)
}

func extractUserIP(req *http.Request) string {
	IPArray := req.Header["X-Forwarded-For"]
	backupIPArray := req.Header["X-Appengine-Remote-Addr"]
	if len(IPArray) < 1 {
		if len(backupIPArray) < 1 {
			return ""
		}
		return backupIPArray[0]
	} else {
		IPWithPotentialProxies := IPArray[0]
		primaryIP := strings.Split(IPWithPotentialProxies, ",")
		return primaryIP[0]
	}
}

func getUserInfo(req *http.Request) (*queries.UserExistsQueryRequest, error) {
	var userInfo queries.UserExistsQueryRequest
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&userInfo)
	if err != nil {
		return nil, err
	} else {
		return &userInfo, nil
	}

}

func generateCookie(userIP string) (*http.Cookie, error) {
	expiration := time.Now().AddDate(0, 0, 15)
	formatedExpiration := expiration.Format(time.RFC3339)
	signature, err := generateSignature(userIP, formatedExpiration)
	if err != nil {
		return nil, err
	}
	val := serializeCookieVals(&types.RegistrationCookie{
		UserIP:     userIP,
		Expiration: formatedExpiration,
		Signature:  signature,
	})
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     RegCookieName,
		Value:    val,
		Expires:  expiration,
		Secure:   true,
		HttpOnly: true,
		Domain:   os.Getenv("DOMAIN"),
		SameSite: http.SameSiteStrictMode,
	}
	return cookie, nil
}

func checkForRegistrationCookie(res http.ResponseWriter, req *http.Request) (*http.Cookie, bool) {
	cookie, err := req.Cookie(RegCookieName)
	if err != nil {
		return nil, false
	}
	return cookie, true
}

func validateRegistrationCookie(cookie *http.Cookie, userIP string) bool {
	val := deserializeCookieVals(cookie.Value)
	expiration, err := time.Parse(time.RFC3339, val.Expiration)
	if err != nil {
		println(err.Error())
		return false
	}
	if time.Now().After(expiration) {
		return false
	}
	err = validateSignature(userIP, val.Expiration, val.Signature)
	return err == nil
}

func addNewUserToDatabase(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	userInfo, err := getUserInfo(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)

		return
	}
	userExists, err := queries.DoesUserExist(ctx, userInfo.Username)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)

		return
	}
	if userExists {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	password, err := crypto.HashPass(userInfo.Password)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	listKey, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User_List", nil), &types.User_List{Items: []string{}})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	listKeyVal := listKey.Encode()
	defaultListMap := make(map[string]string)
	defaultListMap["default"] = listKeyVal
	marshaledDefaultListMap, err := json.Marshal(defaultListMap)
	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := types.UserRecord{Username: userInfo.Username, Password: password, List_Array: marshaledDefaultListMap}
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User_Record", nil), &data)
	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)

}

func concatinateCookieString(userIP string, expiration string) string {
	return userIP + expiration + os.Getenv("SERVER_SIG")
}

func validateSignature(userIP string, expiration string, userSignature string) error {
	correctSignature, err := generateSignature(userIP, expiration)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(correctSignature), []byte(concatinateCookieString(userIP, expiration)))
}
func generateSignature(userIP string, expiration string) (string, error) {
	return crypto.HashPass(concatinateCookieString(userIP, expiration))
}

func serializeCookieVals(source *types.RegistrationCookie) string {
	return source.UserIP + "|" + source.Expiration + "|" + source.Signature
}

func deserializeCookieVals(source string) *types.RegistrationCookie {
	values := strings.Split(source, "|")
	return &types.RegistrationCookie{
		UserIP:     values[0],
		Expiration: values[1],
		Signature:  values[2],
	}
}
