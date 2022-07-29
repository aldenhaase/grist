package endpoints

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetUserList(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	cookie := req.Cookies()
	if len(cookie) == 0 {
		res.WriteHeader(http.StatusUnauthorized)
		encoder.Encode("Ya got no cookies")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(cookie[0].Value), []byte(cookie[0].Name+"secret Key"))
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(err.Error())
		return
	}
	encoder.Encode(map[string]interface{}{"listName": "MyList", "listItems": "hello"})
}
