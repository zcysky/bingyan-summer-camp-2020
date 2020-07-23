package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type QQConfig struct {
	QQID uint `json:"qqid"`
}

type ConnectionConfig struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	AuthKey string `json:"auth_key"`
}

type DatabaseConfig struct {
	DatabaseAddress      string `json:"database_address"`
	DatabaseName         string `json:"database_name"`
	CollectionFriendName string `json:"collection_friend_name"`
	CollectionMemoName   string `json:"collection_memo_name"`
	SleepTime            int    `json:"sleep_time"`
}

type ConfigObject struct {
	QQ         QQConfig         `json:"qq"`
	Connection ConnectionConfig `json:"connection"`
	Database   DatabaseConfig   `json:"database"`
}

var Config ConfigObject

func init() {
	file, err := os.Open("./src/config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteStream, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
