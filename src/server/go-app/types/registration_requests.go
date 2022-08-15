package types

import (
	"time"
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
	Username   string
	Password   string
	List_Array []byte
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
	Last_Modified time.Time
	Items         []string
}
type New_Item struct {
	Item      string
	List_Name string
}
type Delete_List struct {
	Items     []string
	List_Name string
}

type Check_Time struct {
	List_Name     string
	Last_Modified time.Time
	List_Array    []string
}
