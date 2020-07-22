package main

import "qqbot/handler"

func main() {

	//conf, _ := config.ParseConfig()
	handler.StartListen()

	/*
		session, _ :=bot.Auth(conf.BotUrl+"/auth", "miku3939")
		bot.Verify(conf.BotUrl+"/verify", session, conf.BotID)

		msg := bot.GroupMessage{
			Session: session,
			Target:  conf.Target[0],
			Message: []bot.Message{
				{Type: "Plain", Text: "初めまして!\n"},
				{Type: "Plain", Text: "初音ミクです！"},
			},
		}
		//id, _ := bot.SendGroupMessage(conf.BotUrl+"/sendGroupMessage", msg)
		//fmt.Println(id)
	*/
}
