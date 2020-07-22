package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"./config"
	"./controller"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
)

func handleSave(robot *gomirai.Bot, event message.Event) error {

	//fmt.Println(e.Sender.Id, fmt.Sprint(e.Sender.Id))
	tmpE := event
	tmpE.EventId = uint(config.ConfigSet.EventCountConfig.Id)
	config.ConfigSet.EventCountConfig.Id++
	err := controller.HandleAddEvent(fmt.Sprint(event.Sender.Id), tmpE)
	if err != nil {
		return err
	}
	_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("your event has been saved successfully"))
	if err != nil {
		return err
	}
	return nil
}

func handleShowAll(robot *gomirai.Bot, event message.Event) error {
	_, err := robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("your event is shown in the following text"))
	if err != nil {
		return err
	}
	allEvent, err := controller.HandleQuery(fmt.Sprint(event.Sender.Id))
	if err != nil {
		fmt.Println(err)
	}
	for _,allEventOne:=range allEvent {
		allEventJSON, err := json.Marshal(allEventOne)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(allEventJSON))
		_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage(string(allEventJSON)))
		if err != nil {
			return err
		}
	}
	//_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.RichMessage(message.MsgType_Json, string(allEventJSON)))
	//if err != nil {
	//	return err
	//}
	return nil
}

func handleDelete(robot *gomirai.Bot, event message.Event,text string) error {

	Regexp := regexp.MustCompile(`^/delete\s([\d]+)$`)
	params := Regexp.FindStringSubmatch(text)
	//fmt.Println(params)
	if len(params) == 2 {
		err := controller.HandleDeleteEvent(fmt.Sprint(event.Sender.Id), params[1])
		if err != nil {
			if err == mongo.ErrNoDocuments {
				_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("there is no such event"))
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("the event has been deleted"))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

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
								err = handleSave(b, e)
								if err != nil {
									fmt.Println(err)
								}
							}

							//显示用户所有事件  在事件中包含/showMyEvent
							if strings.Contains(QQmessage.Text, "/showMyEvent") {
								err = handleShowAll(b, e)
								if err != nil {
									fmt.Println(err)
								}

							}
							//删除用户事件  在事件中仅有 "/delete eventId"
							if strings.Contains(QQmessage.Text, "/delete") {
								err=handleDelete(b,e,QQmessage.Text)
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
