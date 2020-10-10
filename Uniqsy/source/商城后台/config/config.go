package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type JWTInfo struct {
	Secret			string	`json:"secret"`
	EffectiveTime	int		`json:"effective_time"`
}

type MongoDBInfo struct {
	URL 	string	`json:"url"`
}

type KeyWordInfo struct {
	Limit	int 	`json:"limit"`
}

type ConfigInfo struct {
	JWT 	JWTInfo		`json:"jwt"`
	MongoDB	MongoDBInfo	`json:"mongo_db"`
	KeyWord	KeyWordInfo	`json:"keyword"`
}

var Config ConfigInfo

func Init() {
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}