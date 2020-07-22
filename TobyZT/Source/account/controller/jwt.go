/* This file contains functions to complete jwt verification */

package controller

import (
	"account/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken accepts a LoginForm struct and generate a jwt
func GenerateToken(userid string, admin bool) (tokenStr string, err error) {
	var jwtInfo model.JWTInfo
	err = ParseJson("config/jwt.json", &jwtInfo)
	if err != nil {
		return "", err
	}

	//t1 = now, t2 = now + exp
	t1 := time.Now().Unix()
	t2 := time.Now().Add(time.Minute * time.Duration(jwtInfo.Expire)).Unix()
	claims := jwt.MapClaims{
		"userid": userid,
		"admin":  admin,
		"iat":    t1,
		"nbf":    t1,
		"exp":    t2,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(jwtInfo.Secret))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenStr, nil
}

// ParseToken accepts a jwt string and decrypts it
func ParseToken(tokenStr string) (userid string, admin bool, valid bool, err error) {
	var jwtInfo model.JWTInfo
	err = ParseJson("config/jwt.json", &jwtInfo)
	if err != nil {
		return "", false, false, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtInfo.Secret), nil
	})
	if err != nil {
		return "", false, false, err
	}
	if !token.Valid {
		log.Println("Invalid")
		return "", false, false, nil
	}
	claim, _ := token.Claims.(jwt.MapClaims)
	return claim["userid"].(string), claim["admin"].(bool), true, nil
}

// ParseJson reads a json file from a particular path
// and binds it to JWTInfo variables
func ParseJson(path string, jwtInfo *model.JWTInfo) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, jwtInfo)
	if err != nil {
		return err
	}
	return nil
}
