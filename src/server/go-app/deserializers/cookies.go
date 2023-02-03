package deserializers

import (
	"server/lystrTypes"
	"strings"
)

func SessionCookie(source string) lystrTypes.SessionCookie {
	values := strings.Split(source, "|")
	return lystrTypes.SessionCookie{
		Username:   values[0],
		Expiration: values[1],
		Signature:  values[2],
	}
}
