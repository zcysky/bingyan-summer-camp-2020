package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type MongoConfig struct {
	DBAddress               string `json:"db_address"`
	DBName                  string `json:"db_name"`
	UserCollection          string `json:"user_collection"`
	CourseCollection        string `json:"course_collection"`
	CourseCommentCollection string `json:"course_comment_collection"`
}

type JwtConfig struct {
	JWTTokenLife     int64  `json:"jwt_token_life"`
	JWTSigningMethod string `json:"jwt_signing_method"`
}

type WeiXinConfig struct {
	RequestAddress string `json:"request_address"`
	AppId          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
}

type ConfigObject struct {
	Mongo  MongoConfig  `json:"mongo"`
	Jwt    JwtConfig    `json:"jwt"`
	WeiXin WeiXinConfig `json:"wei_xin"`
}

var Config ConfigObject

func init() {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteStream, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
