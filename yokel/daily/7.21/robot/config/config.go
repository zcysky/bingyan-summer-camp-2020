package config

import (
	"encoding/json"
	"io/ioutil"
)

type mirai struct {
	QQNumber uint `json:"qq-number"`
	ClientHost string `json:"client-host"`
	AuthKey string `json:"auth-key"`
	TargetId uint `json:"target-id"`
}

type eventCount struct {
	Id int64 `json:"id"`
}

type config struct {
	MiraiConfig mirai `json:"mirai"`
	EventCountConfig eventCount `json:"event-count"`
}

const (
	ConfigAddress="./config/config.json"
)

var ConfigSet config

func init(){
	configContent, err := ioutil.ReadFile(ConfigAddress)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configContent, &ConfigSet)
	if err != nil {
		panic(err)
	}

}