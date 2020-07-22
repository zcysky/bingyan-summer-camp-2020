package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BotUrl string `json:"botUrl"`
	AuthKey string `json:"authKey"`
	BotID  int `json:"botID"`
	Target []int  `json:"target"`
}

func ParseConfig() (conf Config, err error) {
	f, err := os.Open("config/config.json")
	if err != nil {
		return conf, err
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(contents, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
