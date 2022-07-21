package queries

import (
	"context"

	"google.golang.org/appengine/v2/datastore"
)

func UserExists(ctx context.Context, username string) (bool, error) {
	query := datastore.NewQuery(username)
	query.Run(ctx)
	num, err := query.Count(ctx)
	if num > 0 {
		return true, err
	} else {
		return false, err
	}

}
