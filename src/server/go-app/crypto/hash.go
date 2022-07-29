package crypto

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	var bytes = []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	return string(hash), err
}
