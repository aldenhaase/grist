package dbFuncs

import (
	"context"
	"log"
	"server/lystrTypes"

	"google.golang.org/appengine/v2/datastore"
)

func GetLists(locs []lystrTypes.ListLocation, ctx context.Context) []lystrTypes.List {

	keys := []*datastore.Key{}
	for _, item := range locs {
		key, _ := datastore.DecodeKey(item.Key)
		keys = append(keys, key)
	}
	arr := make([]lystrTypes.List, len(keys))

	err := datastore.GetMulti(ctx, keys, arr)
	if err != nil {
		log.Println(err.Error())
	}
	return arr
}

func GetUserRecord(username string, ctx context.Context) lystrTypes.UserRecord {
	record := lystrTypes.UserRecord{}
	key := datastore.NewKey(ctx, lystrTypes.User_Record_T, username, 0, nil)
	datastore.Get(ctx, key, &record)
	return record
}
