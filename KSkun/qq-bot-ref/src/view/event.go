package view

import (
	"github.com/Logiase/gomirai/message"
	"qq-bot-ref/config"
	"qq-bot-ref/controller"
	"qq-bot-ref/util"
	"strings"
	"sync"
)

var (
	chanInMapEvent    MsgChanMap // msg in channel for each user in each context
	waitGroupMapEvent WaitGroupMap
)

func initEvent() {
	chanInMapEvent = make(MsgChanMap)
	waitGroupMapEvent = make(WaitGroupMap)
}

func initChan(qq uint) {
	chanInMapEvent[qq] = make(chan string, config.Config.App.ChannelBufferSize)
	waitGroupMapEvent[qq] = &sync.WaitGroup{}
	waitGroupMapEvent[qq].Add(1)
	go deleteChan(qq)
}

func deleteChan(qq uint) {
	waitGroupMapEvent[qq].Wait()
	close(chanInMapEvent[qq])
	delete(chanInMapEvent, qq)
	delete(waitGroupMapEvent, qq)
}

func FilterEvent(msg message.Message, chanOut util.SendMsgChan) {
	if msg.Type != message.MsgType_Plain {
		return
	}
	if chanIn, found := chanInMapEvent[msg.SenderId]; found {
		chanIn <- msg.Text
		return
	}
	// help
	if strings.Index(msg.Text, config.Locale.HelpPrefix) == 0 {
		initChan(msg.SenderId)
		go controller.HandlerHelp(msg.SenderId, msg.Text, chanInMapEvent[msg.SenderId], chanOut, waitGroupMapEvent[msg.SenderId])
	}
	// add
	if strings.Index(msg.Text, config.Locale.AddPrefix+" ") == 0 {
		initChan(msg.SenderId)
		go controller.HandlerAdd(msg.SenderId, msg.Text, chanInMapEvent[msg.SenderId], chanOut, waitGroupMapEvent[msg.SenderId])
	}
	// delete
	if strings.Index(msg.Text, config.Locale.DeletePrefix+" ") == 0 {
		initChan(msg.SenderId)
		go controller.HandlerDelete(msg.SenderId, msg.Text, chanInMapEvent[msg.SenderId], chanOut, waitGroupMapEvent[msg.SenderId])
	}
	// list
	if strings.Index(msg.Text, config.Locale.ListPrefix) == 0 {
		initChan(msg.SenderId)
		go controller.HandlerList(msg.SenderId, msg.Text, chanInMapEvent[msg.SenderId], chanOut, waitGroupMapEvent[msg.SenderId])
	}
}
