package bot

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UpdateReminders() {
	reminders, err := Query()
	if err != nil {
		log.Println(err)
	}
	for _, reminder := range reminders {
		msg := Message{
			Session: session,
			Target:  reminder.ID,
			Message: []MessageChain{
				{Type: "Plain", Text: "【事件提醒】"},
				{Type: "Plain", Text: reminder.Content},
			},
		}
		SendMessage(conf.BotUrl+"/send"+reminder.Type, msg)
	}
}

func SetReminder(msg SingleData, senderID int, i int) {
	s := strings.Split(msg.MessageChain[i].Text, " ")
	if len(s) < 3 {
		reply := Message{Session: session, Target: senderID,
			Message: []MessageChain{{Type: "Plain", Text: "格式错误"}},
		}
		SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
	}
	var advance, gap int
	var err error
	if len(s) == 5 {
		advance, err = strconv.Atoi(s[3])
		if err != nil {
			reply := Message{Session: session, Target: senderID,
				Message: []MessageChain{{Type: "Plain", Text: "格式错误"}},
			}
			SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
		}
		gap, err = strconv.Atoi(s[4])
		if err != nil {
			reply := Message{Session: session, Target: senderID,
				Message: []MessageChain{{Type: "Plain", Text: "格式错误"}},
			}
			SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
		}
	} else {
		advance, gap = 0, 0
	}
	strTime, content := s[1], s[2]
	due, err := parseTime(strTime)
	if err != nil {
		reply := Message{Session: session, Target: senderID,
			Message: []MessageChain{{Type: "Plain", Text: "时间已经过期或格式有误"}},
		}
		SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
		return
	}

	reminder := Reminder{ID: senderID, Type: msg.Type, Due: due.Unix(),
		Content: content, Advance: advance, Gap: gap}
	err = Insert(reminder)
	if err != nil {
		reply := Message{Session: session, Target: senderID,
			Message: []MessageChain{{Type: "Plain", Text: err.Error()}},
		}
		SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
		return
	}
	reply := Message{Session: session, Target: senderID,
		Message: []MessageChain{{Type: "Plain", Text: "添加成功"}},
	}
	SendMessage(conf.BotUrl+"/send"+msg.Type, reply)
}

func QueryReminder(senderID int) {
	reminders, _ := QueryByID(senderID)
	if len(reminders) == 0 {
		return
	}
	var text []MessageChain
	text = append(text, MessageChain{Type: "Plain", Text: "用户" + strconv.Itoa(senderID) + "的备忘录:\n"})
	for _, reminder := range reminders {
		t := time.Unix(reminder.Due, 0).String() + " " + reminder.Content
		text = append(text, MessageChain{Type: "Plain", Text: t + "\n"})
	}
	reply := Message{Session: session, Target: senderID, Message: text}
	SendMessage(conf.BotUrl+"/send"+reminders[0].Type, reply)
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
