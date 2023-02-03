package extractors

import (
	"encoding/json"
	"net/http"
	"server/lystrTypes"
)

func UserFromJSON(req *http.Request) lystrTypes.UserQuery {
	var userInfo lystrTypes.UserQuery
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&userInfo)
	return userInfo
}

//should be generics
func CollaboratorFromJSON(req *http.Request) lystrTypes.CollaboratorQuery {
	var collaborator lystrTypes.CollaboratorQuery
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&collaborator)
	return collaborator
}

func CollectionFromJSON(req *http.Request) lystrTypes.Collection {
	collection := lystrTypes.Collection{}
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&collection)
	return collection
}
func ListFromJSON(req *http.Request) lystrTypes.List {
	list := lystrTypes.List{}
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&list)
	return list
}
func ItemFromJSON(req *http.Request) []lystrTypes.Item {
	items := []lystrTypes.Item{}
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&items)
	return items
}
