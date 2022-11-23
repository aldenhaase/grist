package session

import "strings"

func Unpack(source string) SessionCookie {
	values := strings.Split(source, "|")
	return SessionCookie{
		Username:   values[0],
		Expiration: values[1],
		Signature:  values[2],
	}
}

func Pack(source SessionCookie) string {
	return source.Username + "|" + source.Expiration + "|" + source.Signature
}
