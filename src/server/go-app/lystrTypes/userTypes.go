package lystrTypes

type UserRecord struct {
	Username      string
	Password      string
	ListLocations []ListLocation
}

type UserQuery struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CollaboratorQuery struct {
	ShareWith string `json:"shareWith"`
	ListName  string `json:"listName"`
}
