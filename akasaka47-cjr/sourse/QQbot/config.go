package main

import (
	"encoding/json"
	"fmt"
	io "io/ioutil"
)

type Config struct {
	DataBase   string `json:"database"`
	Collection string `json:"collection"`
}

type Send struct {
	UserId   int64  `json:"user_id"` //发送者QQ号
	NickName string `json:"nickname"`
	Card     string `json:"card"` //群名片／备注
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"` //owner 或 admin 或 member
	Title    string `json:"title"`
}

type Anony struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type PostInfo struct {
	Time        int64  `json:"time"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int    `json:"message_id"`
	GroupId     int64  `json:"group_id"`
	UserId      int64  `json:"user_id"`
	Anonymous   Anony  `json:"anonymous"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	Sender      Send   `json:"sender"`
}

type Respond struct {
	Reply       string `json:"reply"`
	AutoEscape  bool   `json:"auto_escape"`
	AtSender    bool   `json:"at_sender"`
	Delete      bool   `json:"delete"`
	Kick        bool   `json:"kick"`
	Ban         bool   `json:"ban"`
	BanDuration int    `json:"ban_duration"`
}

type Remind struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
