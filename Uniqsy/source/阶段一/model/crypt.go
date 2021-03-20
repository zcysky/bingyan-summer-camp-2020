package model

import (
	"golang.org/x/crypto/bcrypt"
)

//加密密码
func Encrypt(rawStr string) string {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(rawStr),
		bcrypt.DefaultCost)
	return string(hash)
}

//对比密码与数据库中的加密值是否相同
func CheckPasswd(hashed string, raw string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	return err
}