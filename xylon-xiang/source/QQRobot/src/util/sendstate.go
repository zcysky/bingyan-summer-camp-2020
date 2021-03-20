package util

import (
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"log"
)

const (
	SUCCESS = "Set Memorandum successfully!"
	FAILURE = "Set Memorandum failure!"
)

func SendStateMessage(b *gomirai.Bot, userId uint, status string, groupId []uint) {
	if groupId == nil {
		_, err := b.SendFriendMessage(userId, 0, message.PlainMessage(status))
		if err != nil {
			log.Fatal(err)
		}

	} else {
		info := message.Message{
			Target:  userId,
			Display: status,
		}

		_, err := b.SendGroupMessage(groupId[0], 0, info)
		if err != nil {
			log.Fatal(err)
		}
	}
}
