package main

import (
	"encoding/json"
	"io/ioutil"
)
import "fmt"

type Config struct {
	LiveId        []int64 `json:"live_id"`
	FetchInterval int64   `json:"fetch_interval"`
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
