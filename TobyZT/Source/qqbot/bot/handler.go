package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func HandleGroupMessage(ch <-chan SingleData) {
	for {
		msg := <-ch
		fmt.Println(msg)
		matchMessage(msg)
	}

}

func HandleFriendMessage(ch <-chan SingleData) {
	for {
		msg := <-ch
		fmt.Println(msg)
		matchMessage(msg)
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
	for _, data := range chats {
		// chat[0] is target, chat[1] is reply
		chat := strings.Split(data, "--")
		for i := 0; i < len(msg.MessageChain); i++ {
			if len(msg.MessageChain[i].Text) > 7 && msg.MessageChain[i].Text[0:7] == "添加 " {
				status, errMsg := setReminder(msg.MessageChain[i].Text, senderID, msg.Type)
				if status == 1 {
					reply := Message{
						Session: session,
						Target:  senderID,
						Message: []MessageChain{
							{Type: "Plain", Text: errMsg},
						},
					}
					SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
					return
				}
				if status == 0 {
					reply := Message{
						Session: session,
						Target:  senderID,
						Message: []MessageChain{
							{Type: "Plain", Text: "加成功"},
						},
					}
					SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
				}
				return
			}
			if msg.MessageChain[i].Text == chat[0] {
				reply := Message{
					Session: session,
					Target:  senderID,
					Message: []MessageChain{
						{Type: "Plain", Text: chat[1]},
					},
				}
				SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
				return
			}
		}
	}
}

// status 1 - format error
func setReminder(raw string, id int, ty string) (status int, errMsg string) {
	s := strings.Split(raw, " ")
	if len(s) != 3 {
		return 1, "格式错误"
	}
	strTime, content := s[1], s[2]
	due, err := parseTime(strTime)
	if err != nil {
		return 1, "格式错误：时间格式有误"
	}
	reminder := Reminder{
		ID:      id,
		Type:    ty,
		Due:     due.Unix(),
		Content: content,
	}
	err = Insert(reminder)
	if err != nil {
		return 1, "添加失败：" + err.Error()
	}
	if due.Unix()-time.Now().Unix() <= 300 {
		UpdateReminders()
	}
	return 0, ""
}

func parseTime(str string) (due time.Time, err error) {
	ptn := "^([0-9]+).([0-9]+).([0-9]+)(-([0-9]+):([0-9]+))*$"
	reg := regexp.MustCompile(ptn)
	valid := reg.MatchString(str)
	if valid {
		t := reg.FindStringSubmatch(str)
		now := time.Now()
		year, _ := strconv.Atoi(t[1])
		month, _ := strconv.Atoi(t[2])
		day, _ := strconv.Atoi(t[3])
		var hour, min int
		if len(t) == 7 {
			hour, _ = strconv.Atoi(t[5])
			min, _ = strconv.Atoi(t[6])
		} else {
			hour, min = now.Hour(), now.Minute()
		}
		due = time.Date(year, time.Month(month), day,
			hour, min, 0, 0, time.Local)
		if due.Unix() < now.Unix() {
			return due, fmt.Errorf("invalid due")
		}
		return due, nil
	}
	return due, nil
}

func parseReply() (chats []string, err error) {
	f, err := os.Open("config/reply.json")
	if err != nil {
		return chats, err
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return chats, err
	}
	var res Chat
	err = json.Unmarshal(contents, &res)
	if err != nil {
		return chats, err
	}
	return res.Data, nil
}
