package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	io "io/ioutil"
)

type Config struct {
	JWTsecret string `json:"secret"`
	MailName  string `json:"mailname"`
	MainAuth  string `json:"mailauth"`
	MailHost  string `json:"mailhost"`
	MailPort  string `json:"mailport"`
}

//var file_locker sync.Mutex //config file locker

type User struct {
	ID       string `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Checkstr string `json:"checkstr" form:"checkstr" query:"checkstr"`
	Info     string `json:"info" form:"info" query:"info"`
	Status   bool   `json:"status" form:"status" query:"status"`
	Token    string `json:"token" form:"token" query:"token"`
}

type Email struct {
	Name     string `json:"name" form:"name" query:"name"`
	Status   bool   `json:"status" form:"status" query:"status"`
	Info     string `json:"info" form:"info" query:"info"`
	Checkstr string `json:"checkstr" form:"checkstr" query:"checkstr"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}



func InitConfig() bool {
	conf, bl := LoadConfig("config.json") //get config struct
	if !bl {
		fmt.Println("InitConfig failed")
		return false
	}
	Info = conf
	return true
}

func LoadConfig(filename string) (Config, bool) {
	var conf Config
	//file_locker.Lock()
	data, err := io.ReadFile(filename) //read config file
	//file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, false
	}
	datajson := []byte(data)
	//fmt.Println(datajson)
	err = json.Unmarshal(datajson, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	return conf, true
}
