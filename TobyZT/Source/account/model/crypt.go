/* This file contains functions to encrypt and decrypt passwords. */

package model

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(raw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(hash)
}

func Compare(hashed string, raw string) (valid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	if err != nil {
		return false
	}
	return true
}
