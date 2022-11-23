package session

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func getServerKey() string {
	return os.Getenv("SERVER_SIG")
}

func concatinateSignatureValues(cookie SessionCookie) string {
	return cookie.Username + cookie.Expiration + getServerKey()
}

func VerifySignature(cookie SessionCookie) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cookie.Signature), []byte(concatinateSignatureValues(cookie)))
	return err == nil
}
