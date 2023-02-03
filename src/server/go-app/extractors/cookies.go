package extractors

import (
	"net/http"
	"server/deserializers"
	"server/lystrTypes"
)

func ExtractUserSC(req *http.Request) lystrTypes.SessionCookie {
	cookie, _ := req.Cookie(lystrTypes.SCookie_t)
	return deserializers.SessionCookie(cookie.Value)
}
