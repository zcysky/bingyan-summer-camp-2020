package handler

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"qqbot/bot"
	"qqbot/config"
)

var conf config.Config
var session string

func StartListen() {
	conf, _ = config.ParseConfig()
	updateSession()
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 5m", updateSession)
	c.AddFunc("@every 3s", distribute)
	c.Run()
	select {}
}

func updateSession() {
	session, _ = bot.Auth(conf.BotUrl+"/auth", "miku3939")
	bot.Verify(conf.BotUrl+"/verify", session, conf.BotID)
}

func distribute() {
	url := conf.BotUrl + fmt.Sprintf("/fetchMessage?sessionKey=%s&count=10", session)
	buf, _ := bot.Get(url)
	var res bot.FetchResponse
	err := json.Unmarshal(buf, &res)
	if err != nil {
		log.Println(err)
	}
	if len(res.Data) > 0 {
		for i := 0; i < len(res.Data); i++ {
			groupCh := make(chan bot.SingleData)
			friendCh := make(chan bot.SingleData)
			go HandleGroupMessage(groupCh)
			go HandleFriendMessage(friendCh)
			if res.Data[i].Type == "GroupMessage" {
				groupCh <- res.Data[i]
			}
			if res.Data[i].Type == "FriendMessage" {
				friendCh <- res.Data[i]
			}
		}
		fmt.Println(res)

	} else {
		fmt.Println("No data")
	}
}
