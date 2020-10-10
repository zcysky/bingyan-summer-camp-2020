/* This file contains functions to complete jwt verification */

package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken accepts a LoginForm struct and generate a jwt
func GenerateToken(username string) (tokenStr string, err error) {
	var jwtInfo JWTForm
	err = ParseJson("config/jwt.json", &jwtInfo)
	if err != nil {
		return "", err
	}
	//t1 = now, t2 = now + exp
	t1 := time.Now().Unix()
	t2 := time.Now().Add(time.Minute * time.Duration(jwtInfo.Expire)).Unix()
	claims := jwt.MapClaims{
		"username": username,
		"iat":      t1,
		"nbf":      t1,
		"exp":      t2,
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
func ParseToken(tokenStr string) (username string, valid bool, err error) {
	var jwtInfo JWTForm
	err = ParseJson("config/jwt.json", &jwtInfo)
	if err != nil {
		return "", false, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtInfo.Secret), nil
	})
	if err != nil {
		return "", false, err
	}
	if !token.Valid {
		log.Println("Invalid")
		return "", false, nil
	}
	claim := token.Claims.(jwt.MapClaims)
	return claim["username"].(string), true, nil
}

// ParseJson reads a json file from a particular path
// and binds it to JWTInfo variables
func ParseJson(path string, jwtInfo *JWTForm) (err error) {
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
