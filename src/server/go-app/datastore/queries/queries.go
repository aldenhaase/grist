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
	query := datastore.NewQuery("userRecord")
	filt := query.Filter("Username =", username)
	num, err := filt.Count(ctx)
	if num > 0 {
		return true, err
	} else {
		return false, err
	}

}

func DoesPasswordMatch(ctx context.Context, userInfo types.UserRecord) error {
	query := datastore.NewQuery("userRecord")
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
