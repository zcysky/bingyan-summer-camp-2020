package main

import (
	"encoding/json"
	"io/ioutil"
)
import "fmt"

type Config struct {
	LIVEId        []int64 `json:"LIVE_ID"`
	FetchInterval int64   `json:"FetchInterval"`
}

func LoadConfig(file string) Config {
	var config Config
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error when reading", file)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("erro when Unmarshal", configFile)
	}
	return config
}
