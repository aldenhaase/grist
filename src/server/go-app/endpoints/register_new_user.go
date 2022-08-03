package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"server/crypto"
	"server/datastore/queries"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/datastore"
)

func RegisterNewUser(res http.ResponseWriter, req *http.Request) {
	cookie, cookieExists := checkForRegistrationCookie(res, req)
	ipArray := req.Header["X-Forwarded-For"]
	if len(ipArray) < 1 {
		handleLocalhost(res, req)
		return
	}
	userIP := ipArray[0]
	if cookieExists {
		if validateRegistrationCookie(cookie, userIP) {
			addNewUserToDatabase(res, req)
			return
		} else {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		val, err := crypto.HashPass(userIP + "secrete registration key")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookie := &http.Cookie{
			Name:     "LREG",
			Value:    val,
			Expires:  time.Now().AddDate(0, 1, 0),
			Secure:   true,
			HttpOnly: true,
			Domain:   os.Getenv("DOMAIN"),
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
		}
		//res.WriteHeader(http.StatusAccepted)
		http.SetCookie(res, cookie)
		return
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

func handleLocalhost(res http.ResponseWriter, req *http.Request) {
	host := req.Header.Get("Origin")
	if host == "http://localhost:4200" {
		ctx := appengine.NewContext(req)
		userExists, err := queries.DoesUserExist(ctx, "Localhost")
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)

			return
		}
		if userExists {
			res.WriteHeader(http.StatusBadRequest)

			return
		}

		password, _ := crypto.HashPass("d9488dac9bc047ea92d12f6574e27f36967eb751c84e328f50febb636746a8e3")
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "userRecord", nil), &queries.UserExistsQueryRequest{Username: "Localhost", Password: password})
		return
	} else {
		res.WriteHeader(http.StatusBadRequest)

		return
	}
}

func checkForRegistrationCookie(res http.ResponseWriter, req *http.Request) (*http.Cookie, bool) {
	cookie, err := req.Cookie("LYSTRREGISTRATION")
	if err != nil {
		return nil, false
	}
	return cookie, true
}

func validateRegistrationCookie(cookie *http.Cookie, userIP string) bool {
	val := cookie.Value

	err := bcrypt.CompareHashAndPassword([]byte(val), []byte(userIP+"secrete registration key"))
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
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "userRecord", nil), &queries.UserExistsQueryRequest{Username: userInfo.Username, Password: password})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)

}
