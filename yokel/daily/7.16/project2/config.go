package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Secret string `json:"secret"`
	Expire float32 `json:"expireat"`
}
func LoadConfig(file string)(Config,error){
	configContent,err:=ioutil.ReadFile(file)
	if(err!=nil){
		return Config{},err
	}
	var result Config
	err=json.Unmarshal(configContent,&result)
	if(err!=nil){
		return Config{},err
	}
	return result,nil
}
