package util

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	str := ""
	for i := 1; i <= length; i++ {
		str += string(charset[rand.Intn(len(charset))])
	}
	return str
}
