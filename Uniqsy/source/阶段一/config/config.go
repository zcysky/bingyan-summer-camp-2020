package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SMTPStruct struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

type JWTInfo struct {
	Secret			string	`json:"secret"`
	EffectiveTime	int 	`json:"effective_time"`
}

type ConfigStruct struct {
	Invitation 	string		`json:"invitation"`
	SMTP 		SMTPStruct	`json:"smtp"`
	JWT			JWTInfo		`json:"jwt"`
}


var Config ConfigStruct

func Init() {
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}