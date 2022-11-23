package session

type SessionCookie struct {
	Username   string
	Expiration string
	Signature  string
}

type SessionAuthenticationResponse struct {
	Authenticated bool
}
