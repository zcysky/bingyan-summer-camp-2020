package model

import "golang.org/x/crypto/bcrypt"

//加密用户密码
func Encrypt(raw string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//对比用户密码是否是加密后同数据库中密码
func Compare(hashed string, raw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	return err
}
