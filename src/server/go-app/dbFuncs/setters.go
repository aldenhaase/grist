package dbFuncs

import (
	"context"
	"net/http"
	"server/crypto"
	"server/extractors"
	"server/lystrTypes"

	"google.golang.org/appengine/v2/datastore"
)

func incrementIPRecord(ctx context.Context, record lystrTypes.IP_Record, key *datastore.Key) {
	record.Num_Profiles++
	_, err := datastore.Put(ctx, key, &record)
	if err != nil {
		println(err.Error())
	}
}

func CreateUserList(username string, ctx context.Context) (string, error) {
	uuid := username
	key := datastore.NewKey(ctx, lystrTypes.User_List_t, uuid, 0, nil)
	listKey, err := datastore.Put(ctx, key, &lystrTypes.List{Items: []lystrTypes.Item{}, ListName: "default", UUID: uuid})
	if err != nil {
		return "", err
	}
	return listKey.Encode(), nil
}

func SetUserRecord(record lystrTypes.UserRecord, ctx context.Context) {
	key := datastore.NewKey(ctx, lystrTypes.User_Record_T, record.Username, 0, nil)
	datastore.Put(ctx, key, &record)
}

func SetListRecord(list lystrTypes.List, Skey string, ctx context.Context) {
	key, _ := datastore.DecodeKey(Skey)
	datastore.Put(ctx, key, &list)
}

func generateIPRecord(ctx context.Context, record lystrTypes.IP_Record, key *datastore.Key, ipAdress string) {
	record.Num_Profiles = 0
	datastore.Put(ctx, key, &record)
}

func AddNewUserToDatabase(res http.ResponseWriter, req *http.Request, ctx context.Context) {
	userRecordIn := extractors.UserFromJSON(req)

	if DoesUserExist(userRecordIn.Username, ctx) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	password, err := crypto.HashPass(userRecordIn.Password)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	listKey, err := CreateUserList(userRecordIn.Username, ctx)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	userRecordOut := lystrTypes.UserRecord{
		Username:      userRecordIn.Username,
		Password:      password,
		ListLocations: []lystrTypes.ListLocation{{Key: listKey, Deleted: false}},
	}
	_, err = datastore.Put(ctx, datastore.NewKey(ctx, lystrTypes.User_Record_T, userRecordOut.Username, 0, nil), &userRecordOut)
	if err != nil {
		println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
}
