package util

import (
	"2020.7.27/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Encrypt(pwd string) string {
	key := []byte(config.Config.Encrypt.Secret)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(pwd))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
