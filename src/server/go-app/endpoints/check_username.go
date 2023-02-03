package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func CheckUsername(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	encoder := json.NewEncoder(res)
	username := extractors.UserFromJSON(req)

	queryResults := isUsernameAvailable(res, req, ctx, username)
	encoder.Encode(queryResults)

}

func isUsernameAvailable(res http.ResponseWriter, req *http.Request, ctx context.Context, username lystrTypes.UserQuery) bool {
	userExists := dbFuncs.DoesUserExist(username.Username, ctx)
	return !userExists
}
