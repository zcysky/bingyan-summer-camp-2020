package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const VERSION = "1.0"

type TypeAppConfig struct {
	MiraiHost         string `json:"mirai_host"`
	MiraiAuthkey      string `json:"mirai_authkey"`
	QQNumber          uint   `json:"qq_number"`
	TimeLayout        string `json:"time_layout"`
	ChannelBufferSize uint   `json:"channel_buffer_size"`
	Debug             bool   `json:"debug"`
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
