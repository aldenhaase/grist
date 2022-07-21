package queries

import (
	"context"

	"google.golang.org/appengine/v2/datastore"
)

type UserExistsQueryResponse struct {
	Exists bool   `json:"exists"`
	Reason string `json:"reason"`
}

type UserExistsQueryRequest struct {
	Username string `json:"username"`
}

type UserExistsQueryError struct {
	Reason string
}

type User struct {
	Username string `json:"username"`
}

func DoesUserExist(ctx context.Context, username string) (bool, error) {
	query := datastore.NewQuery("")
	query.Run(ctx)
	num, err := query.Count(ctx)
	if num > 0 {
		return true, err
	} else {
		return false, err
	}

}
