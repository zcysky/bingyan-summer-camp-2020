package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DatabaseConfig struct {
	DatabaseAddress    string `json:"database_address"`
	DatabaseName       string `json:"database_name"`
	CollectionUserName string `json:"collection_user_name"`
}

type MysqlConfig struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         string `json:"port"`
}

type JWTConfig struct {
	JWTSecret        string `json:"jwt_secret"`
	JWTSigningMethod string `json:"jwt_signing_method"`
	JWTTokenLife     int64  `json:"jwt_token_life"`
}

type RedisConfig struct {
	RedisAddress      string `json:"redis_address"`
	RedisHistoryLimit int64  `json:"redis_history_limit"`
}

type MailConfig struct {
	MailAddress string `json:"mail_address"`
	SMTPAddress string `json:"smtp_address"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	MailPort    int    `json:"mail_port"`
}

type EncryptConfig struct {
	Secret string `json:"secret"`
}

type ConfigObject struct {
	DataBase DatabaseConfig `json:"data_base"`
	Mysql    MysqlConfig    `json:"mysql"`
	JWT      JWTConfig      `json:"jwt"`
	Redis    RedisConfig    `json:"redis"`
	Mail     MailConfig     `json:"mail"`
	Encrypt  EncryptConfig  `json:"encrypt"`
}

var Config ConfigObject

func init() {
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var config ConfigObject
	err = json.Unmarshal(byteValue, &config)
	Config = config
	if err != nil {
		fmt.Println(err)
	}
}
