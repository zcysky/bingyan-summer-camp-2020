package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	io "io/ioutil"
)

type Config struct {
	JWTsecret   string `json:"secret"`
	Code        string `json:"code"`
	DataBase    string `json:"database"`
	CollectionG string `json:"c-goods"`
	CollectionU string `json:"c-users"`
	CollectionH string `json:"c-hot"`
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

/*-------------------------------------------------------------------------*/
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type HotWord struct {
	Keyword string `json:"keyword"`
	Count   int64  `json:"count"`
	Order   int    `json:"order"`
}

/*-------------------------------------------------------------------------*/

type ReqUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type ReqGoodsList struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Category int    `json:"category"`
	Keyword  string `json:"keyword"`
}

type ReqGood struct {
	Id string `json:"id"`
}

/*-------------------------------------------------------------------------*/
type ResStr struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}

type ResGoodList struct {
	Success bool            `json:"success"`
	Error   string          `json:"error"`
	Data    []*GoodInfoList `json:"data"`
}

type ResHot struct {
	Success bool      `json:"success"`
	Error   string    `json:"error"`
	Data    []*string `json:"data"`
}

type ResGoodInfoDetail struct {
	Success bool            `json:"success"`
	Error   string          `json:"error"`
	Data    *GoodInfoDetail `json:"data"`
}

type ResUserInfo struct {
	Success bool      `json:"success"`
	Error   string    `json:"error"`
	Data    *UserInfo `json:"data"`
}

type ResUserInfoOther struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Data    *UserInfoOther `json:"data"`
}

type ResUserGoods struct {
	Success bool           `json:"success"`
	Error   string         `json:"error"`
	Data    []*GoodInfoSim `json:"data"`
}

type ResPics struct {
	Url string `json:"url"`
}

/*-------------------------------------------------------------------------*/

type UserInfo struct {
	Username          string `json:"username"`
	NickName          string `json:"nickname"`
	Mobile            string `json:"mobile"`
	Email             string `json:"email"`
	TotalViewCount    int    `json:"total_view_count"`
	TotalCollectCount int    `json:"total_collect_count"`
}

type UserInfoAll struct {
	Id                string         `json:"id"`
	Username          string         `json:"username"`
	Password          string         `json:"password"`
	NickName          string         `json:"nickname"`
	Mobile            string         `json:"mobile"`
	Email             string         `json:"email"`
	TotalViewCount    int            `json:"total_view_count"`
	TotalCollectCount int            `json:"total_collect_count"`
	Collections       []*GoodInfoSim `json:"collections"`
}

type UserInfoOther struct {
	NickName          string `json:"nickname"`
	Email             string `json:"email"`
	TotalViewCount    int    `json:"total_view_count"`
	TotalCollectCount int    `json:"total_collect_count"`
}

type GoodInfoList struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Category int    `json:"category"`
	Picture  string `json:"picture"`
}

type GoodInfoALL struct {
	Id           string `json:"id"`
	PubUser      string `json:"pub_user"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Category     int    `json:"category"`
	Price        int    `json:"price"`
	Picture      string `json:"picture"`
	ViewCount    int    `json:"view_count"`
	CollectCount int    `json:"collect_count"`
}

type GoodInfoDetail struct {
	PubUser      string `json:"pub_user"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Category     int    `json:"category"`
	Price        int    `json:"price"`
	Picture      string `json:"picture"`
	ViewCount    int    `json:"view_count"`
	CollectCount int    `json:"collect_count"`
}

type GoodInfoSim struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
