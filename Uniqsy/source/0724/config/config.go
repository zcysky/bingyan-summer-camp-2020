package config

import (
	"encoding/json"
	"io/ioutil"
)

type SMTPStruct struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

type ConfigStruct struct {
	Invitation 	string		`json:"invitation"`
	SMTP 		SMTPStruct	`json:"smtp"`
}


var Config ConfigStruct

func init() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}