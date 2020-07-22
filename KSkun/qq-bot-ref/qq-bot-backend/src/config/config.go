package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TypeAppConfig struct {
	Address     string `json:"address"`
	AccessToken string `json:"access_token"`
}

type TypeMongoConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

type TypeConfig struct {
	App   TypeAppConfig   `json:"app"`
	Mongo TypeMongoConfig `json:"mongo"`
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
