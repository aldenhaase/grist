package auth

import (
	"os"
	"server/crypto"

	"golang.org/x/crypto/bcrypt"
)

func getServerKey() string {
	return os.Getenv("SERVER_SIG")
}

func concatinateSignatureValues(username string, expiration string) string {
	return username + expiration + getServerKey()
}

func VerifySignature(signature string, username string, expiration string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(signature), []byte(concatinateSignatureValues(username, expiration)))
	return err == nil
}

func generateSignature(val string, expiration string) (string, error) {
	return crypto.HashPass(concatinateSignatureValues(val, expiration))
}
