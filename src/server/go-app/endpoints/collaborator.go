package endpoints

import (
	"net/http"
	"server/dbFuncs"
	"server/extractors"

	"google.golang.org/appengine/v2"
)

func Collaborator(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx := appengine.NewContext(req)
	username := extractors.ExtractUserSC(req).Username
	principalRecord := dbFuncs.GetUserRecord(username, ctx)

	collaboratorRequest := extractors.CollaboratorFromJSON(req)
	collaboratorExists := dbFuncs.DoesUserExist(principalRecord.Username, ctx)
	if !collaboratorExists {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	if collaboratorRequest.ShareWith == username {
		res.WriteHeader(http.StatusConflict)
		return
	}
	collaboratorRecord := dbFuncs.GetUserRecord(collaboratorRequest.ShareWith, ctx)

	lists := dbFuncs.GetLists(principalRecord.ListLocations, ctx)
	for index, list := range lists {
		if list.ListName == collaboratorRequest.ListName {
			collaboratorRecord.ListLocations = append(collaboratorRecord.ListLocations, principalRecord.ListLocations[index])
			dbFuncs.SetUserRecord(collaboratorRecord, ctx)
			res.WriteHeader(http.StatusCreated)
			return
		}
	}
}
