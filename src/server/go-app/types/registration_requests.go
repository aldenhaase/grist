package types

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
}
