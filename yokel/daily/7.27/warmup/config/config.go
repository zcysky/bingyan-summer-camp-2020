package config

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type JWTConfig struct {
	Secret        string `json:"secret"`
	TokenDuration int    `json:"token_duration"`
}
type ConfigObject struct {
	JWT JWTConfig `json:"jwt"`
}

type RegisterInfo struct {
	Uid      string `json:"uid",bson:"uid"`
	Pwd      string `json:"pwd",bson:"pwd"`
	Nickname string `json:"nickname",bson:"nickname"`
	Phone    string `json:"phone",bson:"phone"`
	Email    string `json:"email",bson:"email"`
	Type string `json:"type",bson:"type"`
}

type JwtToken struct {
	Uid      string `json:"uid",bson:"uid"`
	Type string `json:"type",bson:"type"`
	jwt.StandardClaims
}

type EmailVelificationCode struct{
	Code string `json:"code"`
}

var Config ConfigObject

const (
	fileAddress = "./config/config.json"
	RedisExpirationTime=3600
)

func init() {
	configContent, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configContent, &Config)
	if err != nil {
		panic(err)
	}
}
