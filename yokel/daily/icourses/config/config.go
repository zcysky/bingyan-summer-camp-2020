package config

import (
	"encoding/json"
	"icourses/defination"
	"io/ioutil"
)

const (
	configJsonAddress="./config/config.json"
)

var Config defination.ConfigObject

func LoadConfig()error{
	configJsonContent,err:=ioutil.ReadFile(configJsonAddress)
	if err!=nil {
		return err
	}
	err=json.Unmarshal(configJsonContent,&Config)
	if err!=nil {
		return err
	}
	return nil
}

func init(){
	err:=LoadConfig()
	if err!=nil{
		panic(err)
	}
}
