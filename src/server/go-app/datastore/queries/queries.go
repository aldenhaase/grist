package queries

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"server/types"
	"sort"
	"time"

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

func GetUserList(username string, ctx context.Context, listName string) (types.User_List, error) {
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
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	key, err := datastore.DecodeKey(listArray[listName])
	if err != nil {
		return list, err
	}
	err = datastore.Get(ctx, key, &list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func SetUserList(username string, ctx context.Context, newItem string, listName string) (types.User_List, error) {
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
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	key, err := datastore.DecodeKey(listArray[listName])
	if err != nil {
		return list, err
	}
	err = datastore.Get(ctx, key, &list)
	if err != nil {
		return types.User_List{}, err
	}

	if len(list.Items) > 100 {
		return types.User_List{}, errors.New("too many items")
	}
	list.Items = append(list.Items, newItem)
	list.Last_Modified = time.Now()
	_, err = datastore.Put(ctx, key, &list)
	return list, err
}
func DeleteListItem(username string, ctx context.Context, itemsToDelete []string, listName string) (types.User_List, error) {
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
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	key, err := datastore.DecodeKey(listArray[listName])
	if err != nil {
		return list, err
	}
	err = datastore.Get(ctx, key, &list)
	if err != nil {
		return types.User_List{}, err
	}

	for _, item := range itemsToDelete {
		if index, res := contains(list.Items, item); res {
			list.Items = append(list.Items[:index], list.Items[index+1:]...)
		}
	}
	list.Last_Modified = time.Now()
	_, err = datastore.Put(ctx, key, &list)
	return list, err
}

func contains(itemsToDelete []string, itemToCheck string) (int, bool) {
	for index, item := range itemsToDelete {
		if item == itemToCheck {
			return index, true
		}
	}
	return -1, false
}

func CreateUserList(ctx context.Context) (string, error) {
	listKey, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User_List", nil), &types.User_List{Items: []string{}, Last_Modified: time.Now()})
	if err != nil {
		return "", err
	}
	return listKey.Encode(), nil
}

func DoesListExist(username string, listName string, ctx context.Context) bool {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return true
	}
	println(len(results))
	if len(results) != 1 {
		return true
	}
	if err != nil {
		return true
	}
	var listArray map[string]string
	err = json.Unmarshal(record[0].List_Array, &listArray)
	if err != nil {
		return true
	}
	if _, contains := listArray[listName]; contains {
		return true
	}
	return false
}

func AddUserList(key string, username string, listName string, ctx context.Context) error {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return err
	}
	println(len(results))
	if len(results) != 1 {
		return errors.New("big Problem")
	}
	if err != nil {
		return errors.New("could not get user record")
	}
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	listArray[listName] = key
	listByteArray, err := json.Marshal(listArray)
	if err != nil {
		return err
	}
	record[0].List_Array = listByteArray
	_, err = datastore.Put(ctx, results[0], &record[0])
	return err
}

func DeleteUserList(username string, listName string, ctx context.Context) error {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return err
	}
	println(len(results))
	if len(results) != 1 {
		return errors.New("big Problem")
	}
	if err != nil {
		return errors.New("could not get user record")
	}
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	listToDelete := listArray[listName]
	keyOfListToDelete, err := datastore.DecodeKey(listToDelete)
	if err != nil {
		return err
	}
	delete(listArray, listName)
	listByteArray, err := json.Marshal(listArray)
	if err != nil {
		return err
	}
	record[0].List_Array = listByteArray
	_, err = datastore.Put(ctx, results[0], &record[0])
	if err != nil {
		return err
	}
	return datastore.Delete(ctx, keyOfListToDelete)
}

func EnumerateLists(username string, ctx context.Context) (map[string]string, error) {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return make(map[string]string), err
	}
	if len(results) != 1 {
		return make(map[string]string), errors.New("big Problem")
	}
	if err != nil {
		return make(map[string]string), errors.New("crypto.hashpass failed")
	}
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)

	return listArray, nil
}

func HasRecordBeenModified(username string, listName string, lastModified time.Time, ctx context.Context) bool {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return false
	}
	if len(results) != 1 {
		return false
	}
	if err != nil {
		return false
	}
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	listToCheck := listArray[listName]
	keyOfListToCheck, err := datastore.DecodeKey(listToCheck)
	if err != nil {
		return false
	}
	var list types.User_List
	err = datastore.Get(ctx, keyOfListToCheck, &list)
	if err != nil {
		return false
	}
	if list.Last_Modified.After(lastModified) {
		return true
	}
	return false
}

func HasListArrayBeenModified(username string, arr []string, ctx context.Context) bool {
	query := datastore.NewQuery("User_Record")
	query = query.Filter("Username =", username)
	record := []types.UserRecord{}
	results, err := query.GetAll(ctx, &record)
	if err != nil {
		return false
	}
	if len(results) != 1 {
		return false
	}
	if err != nil {
		return false
	}
	var listArray map[string]string
	json.Unmarshal(record[0].List_Array, &listArray)
	var listNames []string
	for key := range listArray {
		listNames = append(listNames, key)
	}
	sort.Strings(listNames)
	return !reflect.DeepEqual(listNames, arr)
}
