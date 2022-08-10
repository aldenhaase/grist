package types

import (
	"google.golang.org/appengine/v2/datastore"
)

type CheckUserNameAvailability struct {
	UserName    string `json:"userName"`
	IsAvailable bool   `json:"isAvailable"`
}

type RegisterNewUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Error struct {
	Reason string `json:"reason"`
}

type UserRecord struct {
	Username string
	Password string
	ListID   *datastore.Key
}

type RegistrationCookie struct {
	UserIP     string
	Expiration string
	Signature  string
}
type Authentication_Cookie struct {
	Username   string
	Expiration string
	Signature  string
}

type IP_Record struct {
	Num_Profiles int
}
type User_List struct {
	Title string
	Items []string
}
type New_Item struct {
	Item string
}
