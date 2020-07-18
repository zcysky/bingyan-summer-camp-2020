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
func GenerateToken(form model.TokenForm) (tokenStr string, err error) {
	var jsonInfo model.JsonInfo
	ParseJson("config/config.json", &jsonInfo)

	//t1 = now, t2 = now + exp
	t1 := time.Now().Unix()
	t2 := time.Now().Add(time.Minute * time.Duration(jsonInfo.Expire)).Unix()
	claims := jwt.MapClaims{
		"userid":   form.UserID,
		"email":    form.Email,
		"password": form.Password,
		"iat":      t1,
		"nbf":      t1,
		"exp":      t2,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(jsonInfo.Secret))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenStr, nil
}

// ParseToken accepts a jwt string and decrypts it
func ParseToken(tokenStr string) (form model.TokenForm, valid bool, err error) {
	var jsonInfo model.JsonInfo
	ParseJson("config/config.json", &jsonInfo)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jsonInfo.Secret), nil
	})
	if err != nil {
		return form, false, err
	}
	if !token.Valid {
		log.Println("Invalid")
		return form, false, nil
	}
	claim, _ := token.Claims.(jwt.MapClaims)
	form = model.TokenForm{
		UserID:   claim["userid"].(string),
		Email:    claim["email"].(string),
		Password: claim["password"].(string),
	}
	return form, true, nil
}

// ParseJson reads a json file from a particular path
// and binds it to JsonInfo variables
func ParseJson(path string, jsonInfo *model.JsonInfo) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(contents, jsonInfo)
	if err != nil {
		log.Fatal(err)
	}
}
