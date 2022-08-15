package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"server/datastore/queries"

	"google.golang.org/appengine/v2"
)

func CheckUsername(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	encoder := json.NewEncoder(res)
	username, err := getUsername(req)
	if err != nil {
		encoder.Encode(queries.UserExistsQueryError{Reason: err.Error()})
		return
	}
	queryResults, err := queryUsername(res, req, ctx, *username)
	if err != nil {
		encoder.Encode(queries.UserExistsQueryError{Reason: err.Error()})
	} else {
		encoder.Encode(queryResults)
	}

}

func getUsername(req *http.Request) (*queries.UserExistsQueryRequest, error) {
	var username queries.UserExistsQueryRequest
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&username)
	if err != nil {
		return nil, err
	} else {
		return &username, nil
	}

}

func queryUsername(res http.ResponseWriter, req *http.Request, ctx context.Context, username queries.UserExistsQueryRequest) (*queries.UserExistsQueryResponse, error) {
	userExists, err := queries.DoesUserExist(ctx, username.Username)
	if err != nil {
		return nil, err
	} else {
		if userExists {
			return &queries.UserExistsQueryResponse{Exists: true, Reason: "List Name Already Taken"}, nil
		} else {
			return &queries.UserExistsQueryResponse{Exists: false, Reason: ""}, nil
		}
	}

}
