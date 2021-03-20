package bot

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"qqbot/config"
)

var c *cron.Cron
var conf config.Config
var session string
var groupCh chan SingleData
var friendCh chan SingleData

func StartListen() {
	conf, _ = config.ParseConfig()
	groupCh = make(chan SingleData, 10)
	friendCh = make(chan SingleData, 10)
	go HandleGroupMessage(groupCh)
	go HandleFriendMessage(friendCh)
	updateSession()
	UpdateReminders()
	c = cron.New(cron.WithSeconds())
	c.AddFunc("@every 5m", updateSession)
	c.AddFunc("@every 4s", distribute)
	c.AddFunc("@every 1m", UpdateReminders)
	c.Run()
	select {}
}

func updateSession() {
	session, _ = Auth(conf.BotUrl+"/auth", "miku3939")
	Verify(conf.BotUrl+"/verify", session, conf.BotID)
}

func distribute() {
	url := conf.BotUrl + fmt.Sprintf("/fetchMessage?sessionKey=%s&count=10", session)
	buf, _ := Get(url)
	var res FetchResponse
	err := json.Unmarshal(buf, &res)
	if err != nil {
		log.Println(err)
	}
	if len(res.Data) > 0 {
		for i := 0; i < len(res.Data); i++ {
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
