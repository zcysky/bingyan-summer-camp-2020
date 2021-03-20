package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"

	"github.com/sirupsen/logrus"

	"./config"
	"./controller"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
)


func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := gomirai.NewClient("default", config.ConfigSet.MiraiConfig.ClientHost, config.ConfigSet.MiraiConfig.AuthKey)
	c.Logger.Level = logrus.TraceLevel
	key, err := c.Auth()
	if err != nil {
		c.Logger.Fatal(err)
	}
	b, err := c.Verify(config.ConfigSet.MiraiConfig.QQNumber, key)
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

	for {
		select {
		case e := <-b.Chan:
			switch e.Type {
			case message.EventReceiveFriendMessage:
				if e.Sender.Id == config.ConfigSet.MiraiConfig.TargetId {

					for _, QQmessage := range e.MessageChain {
						//命令式事件处理
						switch QQmessage.Type {
						case message.MsgType_Plain:

							//保存事件  在事件中包含/save
							if strings.Contains(QQmessage.Text, "/save") {
								err = controller.HandleSave(b, e)
								if err != nil {
									fmt.Println(err)
								}
							}

							//显示用户所有事件  在事件中包含/showMyEvent
							if strings.Contains(QQmessage.Text, "/showMyEvent") {
								err = controller.HandleShowAll(b, e)
								if err != nil {
									fmt.Println(err)
								}

							}
							//删除用户事件  在事件中仅有 "/delete eventId"
							if strings.Contains(QQmessage.Text, "/delete") {
								err=controller.HandleDelete(b,e,QQmessage.Text)
								if err != nil {
									fmt.Println(err)
								}

							}
							if strings.Contains(QQmessage.Text, "/addNoti") {
								err=controller.HandleAddNoti(b,e,QQmessage.Text)
								if err != nil {
									fmt.Println(err)
								}
							}
						}
					}

					//非命令式消息回复
					for _, QQmessage := range e.MessageChain {
						switch QQmessage.Type {
						case message.MsgType_Plain:
							switch QQmessage.Text {
							case "call robot":
								_, err = b.SendFriendMessage(e.Sender.Id, 0, message.PlainMessage("hello,this is robot"))
								if err != nil {
									fmt.Println(err)
								}
							case "hello":
								_, err = b.SendFriendMessage(e.Sender.Id, 0, message.PlainMessage("hello,this is robot"))
								if err != nil {
									fmt.Println(err)
								}
								_, err = b.SendFriendMessage(e.Sender.Id, 0, message.PlainMessage("what can i do for you"))
								if err != nil {
									fmt.Println(err)
								}
							case "image":
								var newMessage []message.Message
								newMessage = append(newMessage, message.ImageMessage("path", "image.jpg"))
								newMessage = append(newMessage, message.PlainMessage("your image"))
								fmt.Println(newMessage)
								_, err = b.SendFriendMessage(e.Sender.Id, 0, newMessage...)

								if err != nil {
									fmt.Println(err)
								}
								//_, err = b.SendFriendMessage(e.Sender.Id, 0,message.ImageMessage("id","123"))
								//if err != nil {
								//	fmt.Println(err)
								//}
								//_, err = b.SendFriendMessage(e.Sender.Id, 0,message.AtMessage(e.Sender.Id))
								//if err != nil {
								//	fmt.Println(err)
								//}
							}
						}
					}
				}
			}




		case <-interrupt:
			fmt.Println("######")
			fmt.Println("interrupt")
			fmt.Println("######")
			//c.Release(qq)
			saveFile, err := json.Marshal(config.ConfigSet)
			if err != nil {
				fmt.Println(err)
			}
			ioutil.WriteFile(config.ConfigAddress, saveFile, 0644)
			c.Release(config.ConfigSet.MiraiConfig.QQNumber)
			return
		}

	}
}
