package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	return string(bytes), err
}