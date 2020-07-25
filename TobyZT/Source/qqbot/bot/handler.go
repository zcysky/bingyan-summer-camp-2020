package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func HandleGroupMessage(ch <-chan SingleData) {
	for {
		msg := <-ch
		fmt.Println(msg)
		go matchMessage(msg)
	}

}

func HandleFriendMessage(ch <-chan SingleData) {
	for {
		msg := <-ch
		fmt.Println(msg)
		go matchMessage(msg)
	}
}

func matchMessage(msg SingleData) {
	chats, err := parseReply()
	if err != nil {
		log.Println(err)
	}
	var senderID int
	if msg.Type == "FriendMessage" {
		senderID = msg.Sender.ID
	} else {
		senderID = msg.Sender.Group.ID
	}
	for _, chat := range chats {
		for i := 0; i < len(msg.MessageChain); i++ {
			if len(msg.MessageChain[i].Text) > 7 && msg.MessageChain[i].Text[0:7] == "添加 " {
				SetReminder(msg, senderID, i)
				return
			}
			if msg.MessageChain[i].Text == "查询备忘录" {
				QueryReminder(msg.Sender.ID)
				return
			}
			if msg.MessageChain[i].Text == chat.Key {
				reply := Message{Session: session, Target: senderID,
					Message: []MessageChain{{Type: "Plain", Text: chat.Value}},
				}
				SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
				return
			}
		}
	}
}

func parseReply() (chats []Chat, err error) {
	f, err := os.Open("config/reply.json")
	if err != nil {
		return chats, err
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return chats, err
	}
	var res []Chat
	err = json.Unmarshal(contents, &res)
	if err != nil {
		return chats, err
	}
	return res, nil
}
