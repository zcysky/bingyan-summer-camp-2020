package main

import (
	"QQRobot/src/config"
	"QQRobot/src/memo"
	"fmt"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	qq := config.Config.QQ.QQID

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := gomirai.NewClient(config.Config.Connection.Name,
		config.Config.Connection.Url, config.Config.Connection.AuthKey)

	c.Logger.Level = logrus.TraceLevel
	key, err := c.Auth()
	if err != nil {
		c.Logger.Fatal(err)
	}
	b, err := c.Verify(qq, key)
	if err != nil {
		c.Logger.Fatal(err)
	}
	//defer c.Release(qq)

	go func() {
		err = b.FetchMessages()
		if err != nil {
			c.Logger.Fatal(err)
		}
	}()

	// remind function
	go memo.Reminder(b)

	for {
		select {
		case e := <-b.Chan:
			switch e.Type {
			case message.EventReceiveFriendMessage:
				contexts := e.MessageChain

				for _, context := range contexts {

					// the memorandum function
					memo.Memorandum(b, e.Sender.Id, context.Text)

					switch context.Text {
					case "hello":
						_, err = b.SendFriendMessage(e.Sender.Id, 0,
							message.PlainMessage("Hello "+e.Sender.MemberName))

					case "功能":
						_, err = b.SendFriendMessage(e.Sender.Id, 0,
							message.PlainMessage(" -memorandum function \n -hello function"))

					case "备忘录":
						// send the picture showing how the notes function works
						_, err = b.SendFriendMessage(e.Sender.Id, 0,
							message.ImageMessage("path", "hair.jpg"))
					}
				}

				_, err = b.SendFriendMessage(e.Sender.Id, 0, message.PlainMessage("complete"))
				if err != nil {
					fmt.Println(err)
				}

			case message.EventReceiveGroupMessage:
				if e.Sender.Group.Id == 383229137 {

					contexts := e.MessageChain

					for _, context := range contexts {

						// the memorandum function
						memo.Memorandum(b, e.Sender.Id, context.Text, e.Sender.Group.Id)

						switch context.Text {
						case "hello":
							_, err = b.SendGroupMessage(e.Sender.Group.Id, 0,
								message.PlainMessage("Hello "+e.Sender.MemberName))

						case "功能":
							_, err = b.SendGroupMessage(e.Sender.Group.Id, 0,
								message.PlainMessage(" -memorandum function \n -hello function"))

						case "备忘录":
							// send the picture showing how the notes function works
							_, err = b.SendGroupMessage(e.Sender.Group.Id, 0,
								message.ImageMessage("path", "hair.jpg"))
						}
					}

					_, err = b.SendGroupMessage(e.Sender.Group.Id, 0, message.PlainMessage("complete"))
					if err != nil {
						fmt.Println(err)
					}
				}

			}
		case <-interrupt:
			fmt.Println("######")
			fmt.Println("interrupt")
			fmt.Println("######")
			c.Release(qq)
			return
		}

	}
}
