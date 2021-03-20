package model

import "golang.org/x/crypto/bcrypt"

func Encrypt(rawStr string) string {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(rawStr),
		bcrypt.DefaultCost)
	return string(hash)
}