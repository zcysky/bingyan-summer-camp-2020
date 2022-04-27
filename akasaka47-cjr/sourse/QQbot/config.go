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

type PostInfo struct {
	Time        int64  `json:"time"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"` //消息子类型，如果是好友则是 friend，如果从群或讨论组来的临时会话则分别是 group、discuss
	MessageId   int    `json:"message_id"`
	GroupId     int64  `json:"group_id"`
	UserId      int64  `json:"user_id"`   //发送者 QQ 号
	Anonymous   Anony  `json:"anonymous"` //匿名信息，如果不是匿名消息则为 nil
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	Sender      Send   `json:"sender"` //发送人信息
}

type Anony struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
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

type Respond struct {
	Reply       string `json:"reply"`
	AutoEscape  bool   `json:"auto_escape"`  //消息内容是否作为纯文本发送（即不解析 CQ 码），只在 reply 字段是字符串时有效
	AtSender    bool   `json:"at_sender"`    //是否要在回复开头 at 发送者（自动添加），发送者是匿名用户时无效
	Delete      bool   `json:"delete"`       //撤回该条消息
	Kick        bool   `json:"kick"`         //把发送者踢出群组（需要登录号权限足够），不拒绝此人后续加群请求，发送者是匿名用户时无效
	Ban         bool   `json:"ban"`          //把发送者禁言 ban_duration 指定时长，对匿名用户也有效
	BanDuration int    `json:"ban_duration"` //禁言时长
}

type RespondType struct {
	Private struct {
		Use        string
		UserId     string
		Message    string
		AutoEscape string
	}
	Group struct {
		Use        string
		GroupId    string
		Message    string
		AutoEscape string
	}
	Discuss struct {
		Use        string
		DiscussId  string
		Message    string
		AutoEscape string
	}
	Delete struct {
		Use       string
		MessageID string
	}
	Kick struct {
		Use           string
		GroupID       string
		UserID        string
		RejectRequest string
	}
	Ban struct {
		Use      string
		GroupID  string
		UserID   string
		Duration string
	}
}

type Remind struct {
	//Type    string `json:"type"`
	ID       string `json:"id"`
	QQ       int64  `json:"qq"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	DDL      int64  `json:"ddl"`
	//Ahead    int    `json:"ahead"`
	//Interval int    `json:"interval"`
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

var Url string

func InitType() {
	Url = "http://192.168.31.6:5700/"
	T.Private.Use = "send_private_msg?" //发送私聊消息
	T.Private.UserId = "user_id="
	T.Private.Message = "&&message="
	T.Private.AutoEscape = "&&auto_escape=" //消息内容是否作为纯文本发送（即不解析 CQ 码），只在 message 字段是字符串时有效
	T.Group.Use = "send_group_msg?"         // 发送群消息
	T.Group.GroupId = "group_id="
	T.Group.Message = "&&message="
	T.Group.AutoEscape = "&&auto_escape="
	T.Discuss.Use = "send_discuss_msg?" //发送讨论组消息
	T.Discuss.DiscussId = "discuss_id="
	T.Discuss.Message = "&&message="
	T.Discuss.AutoEscape = "&&auto_escape="
	T.Delete.Use = "delete_msg?" //撤回消息
	T.Delete.MessageID = "message_id="
	T.Kick.Use = "set_group_kick?" //群组踢人
	T.Kick.GroupID = "group_id="
	T.Kick.UserID = "&&user_id="
	T.Kick.RejectRequest = "&&reject_add_request=" //拒绝此人的加群请求
	T.Ban.Use = "set_group_ban?"                   //群组单人禁言
	T.Ban.GroupID = "group_id="
	T.Ban.UserID = "&&user_id="
	T.Ban.Duration = "&&duration="
}
