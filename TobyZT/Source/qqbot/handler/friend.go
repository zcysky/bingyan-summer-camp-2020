package handler

import "qqbot/bot"

func HandleFriendMessage(ch <-chan bot.SingleData) {
	msg := <-ch
	matchFriendMessage(msg, "ping", "pong")
}

func matchFriendMessage(msg bot.SingleData, target string, reply string) {
	senderID := msg.Sender.ID
	if msg.MessageChain[1].Text == target {
		reply := bot.GroupMessage{
			Session: session,
			Target:  senderID,
			Message: []bot.MessageChain{
				{Type: "Plain", Text: reply},
			},
		}
		bot.SendMessage(conf.BotUrl+"/sendFriendMessage", reply)
	}
}
