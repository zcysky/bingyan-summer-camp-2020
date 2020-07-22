package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TypeAppConfig struct {
	Address          string `json:"address"`
	VerifyCodeLength int    `json:"verify_code_length"`
	VerifyCodeExpire int    `json:"verify_code_expire"` // unit: minute
}

type TypeJWTConfig struct {
	TokenExpire int    `json:"token_expire"` // unit: minute
	Secret      string `json:"secret"`
}

type TypeMongoConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

type TypeRedisConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type TypeConfig struct {
	App   TypeAppConfig   `json:"app"`
	JWT   TypeJWTConfig   `json:"jwt"`
	Mongo TypeMongoConfig `json:"mongo"`
	Redis TypeRedisConfig `json:"redis"`
}

var Config TypeConfig

func InitConfig() {
	configFilename := "default.json" // use `default.json` as default filename
	// set env variable CONFIG_FILE to use other config file
	if filename, ok := os.LookupEnv("CONFIG_FILE"); ok {
		configFilename = filename
	}

	configFile, err := ioutil.ReadFile("./config/" + configFilename)
	if err != nil {
		log.Println("config: error when read config file " + configFilename)
		log.Panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Println("config: error when unmarshal config")
		log.Panic(err)
	}

	log.Println("config: config " + configFilename + " loaded")
}
