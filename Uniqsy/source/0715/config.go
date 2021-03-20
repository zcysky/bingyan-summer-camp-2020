package main

import (
	"encoding/json"
	"io/ioutil"
)

type RoomConfig struct {
	Roomid	int	`json:"room_id"`
}

type ConfigStruct struct {
	Cookie		string     `json:"cookie"`
	UserAgent	string      `json:"user_agent"`
	Host 		string       `json:"host"`
	Interval	int          `json:"interval"`
	Count		int         `json:"count"`
	List		[]RoomConfig `json:"list"`
}

var Config ConfigStruct

func Init() {
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}