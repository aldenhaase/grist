package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
)

func ListGrabber(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")

	ctx := appengine.NewContext(req)
	username := extractors.ExtractUserSC(req).Username
	userRecord := dbFuncs.GetUserRecord(username, ctx)
	lists := dbFuncs.GetLists(userRecord.ListLocations, ctx)
	log.Println(lists)
	collection := createCollection(lists, userRecord.ListLocations)
	encoder.Encode(collection)
}

func createCollection(lists []lystrTypes.List, keys []lystrTypes.ListLocation) lystrTypes.Collection {
	collection := lystrTypes.Collection{}
	for index, list := range lists {
		if !keys[index].Deleted {
			redactedList := []lystrTypes.Item{}
			for _, item := range list.Items {
				if !itemIsDeleted(item, list) {
					redactedList = append(redactedList, item)
				}
			}
			list.Items = redactedList
			collection.Lists = append(collection.Lists, list)
		}
	}
	return collection
}

func itemIsDeleted(item lystrTypes.Item, list lystrTypes.List) bool {
	for _, delUUID := range list.DeletedItems {
		if item.UUID == delUUID {
			return true
		}
	}
	return false
}
