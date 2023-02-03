package endpoints

import (
	"context"
	"net/http"
	"server/dbFuncs"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/datastore"
)

func ListSetter(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx := appengine.NewContext(req)
	username := extractors.ExtractUserSC(req).Username
	setList(username, req, ctx)
}

func setList(username string, req *http.Request, ctx context.Context) {
	record := dbFuncs.GetUserRecord(username, ctx)
	incomingCollection := extractors.CollectionFromJSON(req)
	existingCollection := dbFuncs.GetLists(record.ListLocations, ctx)
	//if incoming collection does not have list but existing collection does have list
	//
	addNewItemsToRecord(incomingCollection, &record.ListLocations, existingCollection, ctx)
	dbFuncs.SetUserRecord(record, ctx)
}

func addNewItemsToRecord(incomingCollection lystrTypes.Collection, existingListLocs *[]lystrTypes.ListLocation, existingCollection []lystrTypes.List, ctx context.Context) {
	for _, incomingList := range incomingCollection.Lists {
		key := datastore.NewKey(ctx, lystrTypes.User_List_t, incomingList.UUID, 0, nil)
		existsAt, exists := doesListExist(incomingList, existingCollection, ctx)
		if exists {
			existingList := existingCollection[existsAt]
			incomingList.DeletedItems = existingList.DeletedItems
			missingItems := getMissingItems(incomingList, existingList, ctx)
			if len(missingItems) > 0 {
				incomingList.Items = append(incomingList.Items, missingItems...)
			}
			datastore.Put(ctx, key, &incomingList)
		} else {
			datastore.Put(ctx, key, &incomingList)
			(*existingListLocs) = append((*existingListLocs), lystrTypes.ListLocation{Key: key.Encode(), Deleted: false})
		}
	}
}
func doesListExist(incomingList lystrTypes.List, existingLists []lystrTypes.List, ctx context.Context) (int, bool) {
	for index, existingList := range existingLists {
		if existingList.UUID == incomingList.UUID {
			return index, true
		}
	}
	return -1, false
}
func getMissingItems(incomingList lystrTypes.List, existingList lystrTypes.List, ctx context.Context) []lystrTypes.Item {
	missingItems := []lystrTypes.Item{}
	for _, existingItem := range existingList.Items {
		missing := true
		for _, incomingItem := range incomingList.Items {
			if incomingItem.UUID == existingItem.UUID {
				missing = false
				break
			}
		}
		if missing {
			missingItems = append(missingItems, existingItem)
		}
	}
	return missingItems
}
