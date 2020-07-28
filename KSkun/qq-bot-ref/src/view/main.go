package view

import (
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"log"
	"os"
	"os/signal"
	"qq-bot-ref/config"
	"qq-bot-ref/controller"
	"qq-bot-ref/util"
	"strconv"
	"sync"
)

type MsgChanMap map[uint]chan string
type WaitGroupMap map[uint]*sync.WaitGroup

var Bot *gomirai.Bot

func InitView() {
	initEvent()

	client := gomirai.NewClient("default", config.Config.App.MiraiHost, config.Config.App.MiraiAuthkey)
	client.Logger.Logger.Out = util.LogWriter
	session, err := client.Auth()
	if err != nil {
		log.Println("view: error when mirai auth")
		log.Panic(err)
	}
	Bot, err = client.Verify(config.Config.App.QQNumber, session)
	if err != nil {
		log.Println("view: error when mirai verify")
		log.Panic(err)
	}

	// message sender goroutine
	sendMsgChan := make(util.SendMsgChan, config.Config.App.ChannelBufferSize)
	go func() {
		for msg := range sendMsgChan {
			_, err := Bot.SendFriendMessage(msg.QQ, 0, message.PlainMessage(msg.Msg))
			if err != nil {
				log.Println("view: error when send message to " + strconv.Itoa(int(msg.QQ)))
				log.Println(err)
			}
		}
	}()
	go controller.CheckerRemind(sendMsgChan)

	// message fetcher goroutine
	go func() {
		err = Bot.FetchMessages()
		if err != nil {
			log.Println("view: error when mirai fetch messages")
			log.Panic(err)
		}
	}()

	// message handler goroutine
	msgChan := make(chan message.Message, config.Config.App.ChannelBufferSize)
	go func() {
		for msg := range msgChan {
			if !FilterEvent(msg, sendMsgChan) {
				sendMsgChan <- util.DefaultMsg(msg.SenderId, config.Locale.UnknownOperation)
			}
		}
	}()

	// release session after interrupt
	chanInterrupt := make(chan os.Signal, 1)
	signal.Notify(chanInterrupt, os.Interrupt)

	// split to single message, pass them to handler
	for {
		select {
		case e := <-Bot.Chan:
			switch e.Type {
			case message.EventReceiveFriendMessage:
				for i, msg := range e.MessageChain {
					if i == 0 {
						continue
					}
					msg.SenderId = e.Sender.Id
					msgChan <- msg
				}
			}
		case <-chanInterrupt:
			log.Println("view: INTERRUPT")
			client.Release(config.Config.App.QQNumber)
			return
		}
	}
}
