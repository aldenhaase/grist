package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/datastore/queries"
	"server/validate"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/datastore"
)

func Test(num int) int {
	return num
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		headers := req.Header.Get("Access-Control-Request-Headers")
		(res).Header().Set("Access-Control-Allow-Headers", headers)
		(res).Header().Set("Access-Control-Allow-Origin", "*")
		if req.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(res, req)
	})
}

func limitBodySize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//check what type of request and set body size
		req.Body = http.MaxBytesReader(res, req.Body, 200)
		next.ServeHTTP(res, req)
	})
}

func registerNewUser(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	encoder := json.NewEncoder(res)
	userInfo, err := getUserInfo(req)
	if err != nil {
		encoder.Encode(queries.UserExistsQueryError{Reason: err.Error()})
		return
	}
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "userRecord", nil), &queries.UserExistsQueryRequest{Username: userInfo.Username, Password: userInfo.Password})
	if err != nil {
		encoder.Encode(queries.RegisterUserResponse{Status: 300, Error: err.Error()})
		return
	}
	encoder.Encode(queries.RegisterUserResponse{Status: 0, Error: ""})
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

func queryUsername(res http.ResponseWriter, req *http.Request, ctx context.Context, username queries.UserExistsQueryRequest) (*queries.UserExistsQueryResponse, error) {
	userExists, err := queries.DoesUserExist(ctx, username.Username)
	if err != nil {
		return nil, err
	} else {
		if userExists {
			return &queries.UserExistsQueryResponse{Exists: true, Reason: "username unavailable"}, nil
		} else {
			return &queries.UserExistsQueryResponse{Exists: false, Reason: ""}, nil
		}
	}

}

func checkUsername(res http.ResponseWriter, req *http.Request) {
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

func root(res http.ResponseWriter, req *http.Request) {

}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/checkUsername", checkUsername)
	mux.HandleFunc("/registerNewUser", registerNewUser)
	mux.HandleFunc("/", root)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	mux := http.NewServeMux()
	setupHandlers(mux)
	http.Handle("/", enableCors(limitBodySize(validate.Json((mux)))))
	appengine.Main()
}
