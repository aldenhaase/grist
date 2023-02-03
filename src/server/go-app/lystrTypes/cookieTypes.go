package lystrTypes

type SessionCookie struct {
	Username   string
	Expiration string
	Signature  string
}

var RCookie_t = "LREG"
var SCookie_t = "LAUTH"
