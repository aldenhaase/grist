package endpoints

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GetUserList(res http.ResponseWriter, req *http.Request) {
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
		encoder.Encode(err.Error())
		return
	}
	user := cookie.Username
	encoder.Encode(map[string]interface{}{"listName": "MyList", "listItems": "hello " + user})
}
