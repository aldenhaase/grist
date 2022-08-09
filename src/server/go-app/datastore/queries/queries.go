package queries

import (
	"context"
	"errors"
	"server/types"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/v2/datastore"
)

type UserExistsQueryResponse struct {
	Exists bool   `json:"exists"`
	Reason string `json:"reason"`
}

type UserExistsQueryRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserExistsQueryError struct {
	Reason string
}

type RegisterUserResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type User struct {
	Username string `json:"username"`
}

func DoesUserExist(ctx context.Context, username string) (bool, error) {
	query := datastore.NewQuery("User_Record")
	filt := query.Filter("Username =", username)
	num, err := filt.Count(ctx)
	if num > 0 {
		return true, err
	} else {
		return false, err
	}

}

func DoesPasswordMatch(ctx context.Context, userInfo types.UserRecord) error {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", userInfo.Username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errors.New("big Problem")
	}
	//record := types.UserRecord{}
	if err != nil {
		return errors.New("crypto.hashpass failed")
	}
	return bcrypt.CompareHashAndPassword([]byte(record[0].Password), []byte(userInfo.Password))

}

func HasIpMetQuota(ctx context.Context, ipAdress string) bool {
	record := types.IP_Record{}
	key := datastore.NewKey(ctx, "IP_Record", ipAdress, 0, nil)
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

func generateIPRecord(ctx context.Context, record types.IP_Record, key *datastore.Key, ipAdress string) {
	record.Num_Profiles = 0
	_, err := datastore.Put(ctx, key, &record)
	if err != nil {
		println(err.Error())
	}
}

func incrementIPRecord(ctx context.Context, record types.IP_Record, key *datastore.Key) {
	record.Num_Profiles++
	_, err := datastore.Put(ctx, key, &record)
	if err != nil {
		println(err.Error())
	}
}

func GetUserList(username string, ctx context.Context) (types.User_List, error) {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return types.User_List{}, err
	}
	if len(results) != 1 {
		return types.User_List{}, errors.New("big Problem")
	}
	if err != nil {
		return types.User_List{}, errors.New("crypto.hashpass failed")
	}
	list := types.User_List{}
	key := record[0].ListID
	err = datastore.Get(ctx, key, &list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func SetUserList(username string, ctx context.Context, listRecord types.User_List) error {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errors.New("big Problem")
	}
	if err != nil {
		return errors.New("crypto.hashpass failed")
	}
	key := record[0].ListID
	list := listRecord.Items
	if len(list) > 100 {
		return errors.New("too many items")
	}
	_, err = datastore.Put(ctx, key, &listRecord)
	return err
}
