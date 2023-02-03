package dbFuncs

import (
	"context"
	"server/lystrTypes"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2/datastore"
)

func DoesUserExist(username string, ctx context.Context) bool {
	query := datastore.NewQuery(lystrTypes.User_Record_T).KeysOnly()
	keys, _ := query.GetAll(ctx, nil)
	for _, key := range keys {
		if key.StringID() == username {
			return true
		}
	}
	return false
}

func DoesPasswordMatch(ctx context.Context, userInfo lystrTypes.UserQuery) error {
	record := lystrTypes.UserRecord{}
	key := datastore.NewKey(ctx, lystrTypes.User_Record_T, userInfo.Username, 0, nil)
	datastore.Get(ctx, key, &record)

	return bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(userInfo.Password))

}

func HasIpMetQuota(ctx context.Context, ipAdress string) bool {
	record := lystrTypes.IP_Record{}
	key := datastore.NewKey(ctx, lystrTypes.IP_Record_t, ipAdress, 0, nil)
	err := datastore.Get(ctx, key, &record)
	if err == nil {
		if record.Num_Profiles < 5 {
			incrementIPRecord(ctx, record, key)
			return false
		}
		return true
	}
	if err == datastore.ErrNoSuchEntity {
		generateIPRecord(ctx, record, key, ipAdress)
		return false
	}
	return true
}
