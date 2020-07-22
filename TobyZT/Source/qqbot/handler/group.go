package handler

import "qqbot/bot"

func HandleGroupMessage(ch <-chan bot.SingleData) {
	msg := <-ch
	MatchGroupMessage(msg, "在吗", "wdnmd不在")
	MatchGroupMessage(msg, "来点涩图", "没有，滚")
}

func MatchGroupMessage(msg bot.SingleData, target string, reply string) {
	senderID := msg.Sender.Group.ID
	if msg.MessageChain[1].Text == target {
		reply := bot.GroupMessage{
			Session: session,
			Target:  senderID,
			Message: []bot.MessageChain{
				{Type: "Plain", Text: reply},
			},
		}
		bot.SendMessage(conf.BotUrl+"/sendGroupMessage", reply)
	}
}
