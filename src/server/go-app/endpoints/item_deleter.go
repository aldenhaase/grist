package endpoints

import (
	"net/http"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func ItemDeleter(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	ctx := appengine.NewContext(req)
	username := extractors.ExtractUserSC(req).Username
	itemsToDelete := extractors.ItemFromJSON(req)
	userRecord := dbFuncs.GetUserRecord(username, ctx)
	lists := dbFuncs.GetLists(userRecord.ListLocations, ctx)
	index := getContainingListIndex(itemsToDelete, lists)
	if index != -1 {
		for _, item := range itemsToDelete {
			lists[index].DeletedItems = append(lists[index].DeletedItems, item.UUID)
		}
		dbFuncs.SetListRecord(lists[index], userRecord.ListLocations[index].Key, ctx)
	}
}

func getContainingListIndex(itemsToDelete []lystrTypes.Item, lists []lystrTypes.List) int {
	for index, list := range lists {
		for _, item := range list.Items {
			if item.UUID == itemsToDelete[0].UUID {
				return index
			}
		}
	}
	return -1
}
