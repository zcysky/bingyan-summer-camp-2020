package main

import (
	"JWT/config"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func GenerateToken(user config.User) string {
	var jsonInfo config.JsonInfo
	ParseJson("config/config.json", &jsonInfo)

	//Make claims
	t1 := time.Now().Unix() //t1=now t2=now+exp
	t2 := time.Now().Add(time.Minute * time.Duration(jsonInfo.Expire)).Unix()
	claims := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"iat":      t1,
		"nbf":      t1,
		"exp":      t2,
	}
	// Make token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(jsonInfo.Secret))
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return tokenStr
}

func ParseToken(tokenStr string) jwt.MapClaims {
	var jsonInfo config.JsonInfo
	ParseJson("config/config.json", &jsonInfo)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jsonInfo.Secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if token.Valid {
		log.Println("Valid")
		claim, _ := token.Claims.(jwt.MapClaims)
		return claim
	} else {
		log.Fatal("Invalid!!")
		return nil
	}
}

func ParseJson(path string, jsonInfo *config.JsonInfo) {
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
