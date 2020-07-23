package bot

import (
	"log"
	"time"
)

func UpdateReminders() {
	reminders, err := Query()
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < len(reminders); i++ {
		addReminder(reminders[i])
	}
}

func addReminder(reminder Reminder) {
	go func() {
		sec := time.Duration(reminder.Due - time.Now().Unix())
		time.Sleep(sec * time.Second)

		msg := Message{
			Session: session,
			Target:  reminder.ID,
			Message: []MessageChain{
				{Type: "Plain", Text: "【事件提醒】"},
				{Type: "Plain", Text: reminder.Content},
			},
		}
		SendMessage(conf.BotUrl+"/send"+reminder.Type, msg)
	}()
}
