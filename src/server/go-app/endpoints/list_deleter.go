package endpoints

import (
	"net/http"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func ListDeleter(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx := appengine.NewContext(req)
	username := extractors.ExtractUserSC(req).Username
	listsToDelete := extractors.ListFromJSON(req)
	userRecord := dbFuncs.GetUserRecord(username, ctx)
	lists := dbFuncs.GetLists(userRecord.ListLocations, ctx)
	deleteList(listsToDelete, &userRecord.ListLocations, lists)
	dbFuncs.SetUserRecord(userRecord, ctx)
}

func deleteList(listToDelete lystrTypes.List, listLocs *[]lystrTypes.ListLocation, lists []lystrTypes.List) {
	for listIndex, fromList := range lists {
		if listToDelete.UUID == fromList.UUID {
			(*listLocs)[listIndex].Deleted = true
		}
	}
}
