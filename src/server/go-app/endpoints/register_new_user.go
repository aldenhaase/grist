package endpoints

import (
	"encoding/json"
	"net/http"
	"server/crypto"
	"server/datastore/queries"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/datastore"
)

func RegisterNewUser(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	encoder := json.NewEncoder(res)
	userInfo, err := getUserInfo(req)
	if err != nil {
		encoder.Encode(queries.UserExistsQueryError{Reason: err.Error()})
		return
	}
	userExists, err := queries.DoesUserExist(ctx, userInfo.Username)
	if err != nil {
		encoder.Encode(queries.UserExistsQueryError{Reason: err.Error()})
		return
	}
	if userExists {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("user already exits")
		return
	}
	password, err := crypto.HashPass(userInfo.Password)
	if err != nil {
		encoder.Encode((queries.UserExistsQueryError{Reason: err.Error()}))
		return
	}
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "userRecord", nil), &queries.UserExistsQueryRequest{Username: userInfo.Username, Password: password})
	if err != nil {
		encoder.Encode(queries.RegisterUserResponse{Status: 300, Error: err.Error()})
		return
	}
	encoder.Encode(queries.RegisterUserResponse{Status: 0, Error: ""})
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
