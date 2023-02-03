package lystrTypes

type Item struct {
	Value  string `json:"value"`
	Marked bool   `json:"marked"`
	UUID   string `json:"uuid"`
}

type List struct {
	ListName     string   `json:"listName"`
	UUID         string   `json:"uuid"`
	Items        []Item   `json:"items"`
	DeletedItems []string `json:"-"`
}

type ListLocation struct {
	Key     string
	Deleted bool
}

type Collection struct {
	Lists []List `json:"lists"`
}
