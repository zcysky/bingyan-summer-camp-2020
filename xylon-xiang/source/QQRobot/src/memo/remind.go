package memo

import (
	"QQRobot/src/config"
	"QQRobot/src/module"
	"context"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func Reminder(b *gomirai.Bot) {
	var (
		cursor  *mongo.Cursor
		err     error
		results []bson.M
		now     int64
		filter  bson.D
	)

	for {
		// get the unix timestamp now
		now = time.Now().Unix()

		// find all the memorandums should been send
		filter = bson.D{
			{"remind_time", bson.D{
				{"$lte", now},
				{"$gt", 0},
			},
			}}

		cursor, err = module.MemoCol.Find(context.Background(), filter)
		if err != nil {
			if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
				continue
			}
			log.Fatal(err)
		}

		if err = cursor.All(context.Background(), &results); err != nil {
			log.Fatal(err)
		}

		// send the memorandum
		for _, result := range results {

			_, err = b.SendFriendMessage(uint(result["user_id"].(int64)), 0,
				message.PlainMessage(result["content"].(string)))
			if err != nil {
				log.Fatal(err)
			}

			// when the remind was send, delete the memorandum
			err = deleteMemo(result["memo_id"].(string))
			if err != nil {
				log.Fatal(err)
			}

		}

		time.Sleep(time.Duration(config.Config.Database.SleepTime) * time.Minute)

	}
}
