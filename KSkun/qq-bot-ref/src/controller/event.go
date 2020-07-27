package controller

import (
	"fmt"
	"log"
	"qq-bot-ref/config"
	"qq-bot-ref/model"
	"qq-bot-ref/util"
	"strconv"
	"strings"
	"sync"
	"time"
)

func HandlerHelp(qq uint, msg string, chanIn chan string, chanOut util.SendMsgChan, waitGroup *sync.WaitGroup) {
	chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.Help, config.VERSION))
	waitGroup.Done()
}

func HandlerAdd(qq uint, msg string, chanIn chan string, chanOut util.SendMsgChan, waitGroup *sync.WaitGroup) {
	var event model.Event
	event.User = qq
	event.Desc = strings.Replace(msg, config.Locale.AddPrefix+" ", "", 1)

	// get time
	chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.AddRequireTime, event.Desc, config.Config.App.TimeLayout))
	timeStr := <-chanIn
	timeObj, err := time.ParseInLocation(config.Config.App.TimeLayout, timeStr, time.Local)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncAdd, err.Error())
		return
	}
	event.Time = timeObj.Unix()

	// get if remind needed
	chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.AddRequireRemind, event.Desc))
	remindStr := <-chanIn
	remind, err := util.Boolean(remindStr)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncAdd, err.Error())
		return
	}
	event.Remind = remind

	if remind {
		// get remind before time
		chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.AddRequireRemindBefore, event.Desc))
		remindTimeBeforeStr := <-chanIn
		remindTimeBefore, err := strconv.Atoi(remindTimeBeforeStr)
		if err != nil {
			chanOut <- util.FailedMsg(qq, config.Locale.FuncAdd, err.Error())
			return
		}
		event.RemindTime = timeObj.Add(-time.Duration(remindTimeBefore) * time.Minute).Unix()

		// get remind interval
		chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.AddRequireRemindInterval, event.Desc))
		remindTimeIntervalStr := <-chanIn
		remindTimeInterval, err := strconv.Atoi(remindTimeIntervalStr)
		if err != nil {
			chanOut <- util.FailedMsg(qq, config.Locale.FuncAdd, err.Error())
			return
		}
		event.RemindInterval = remindTimeInterval
	}

	_, err = model.AddEvent(event)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncAdd, err.Error())
		return
	}
	chanOut <- util.SuccessMsg(qq, config.Locale.FuncAdd)
	waitGroup.Done()
}

func HandlerDelete(qq uint, msg string, chanIn chan string, chanOut util.SendMsgChan, waitGroup *sync.WaitGroup) {
	id, err := strconv.Atoi(strings.Replace(msg, config.Locale.DeletePrefix+" ", "", 1))
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncDelete, err.Error())
		return
	}

	events, err := model.GetEventsByUser(qq)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncDelete, err.Error())
		return
	}
	if len(events) < id || id <= 0 {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncDelete, "id out of range")
		return
	}
	event := events[id-1]

	chanOut <- util.DefaultMsg(qq, fmt.Sprintf(config.Locale.DeleteConfirm, id, event.Desc))
	confirmed, err := util.Boolean(<-chanIn)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncDelete, err.Error())
		return
	}
	if !confirmed {
		chanOut <- util.DefaultMsg(qq, config.Locale.DeleteAbort)
		return
	}

	err = model.DeleteEvent(event.ID.Hex())
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncDelete, err.Error())
		return
	}
	chanOut <- util.SuccessMsg(qq, config.Locale.FuncDelete)
	waitGroup.Done()
}

func HandlerList(qq uint, msg string, chanIn chan string, chanOut util.SendMsgChan, waitGroup *sync.WaitGroup) {
	events, err := model.GetEventsByUser(qq)
	if err != nil {
		chanOut <- util.FailedMsg(qq, config.Locale.FuncList, err.Error())
		return
	}
	resp := ""
	for i, event := range events {
		resp += fmt.Sprintf(config.Locale.ListEntry, i+1, event.Desc,
			time.Unix(event.Time, 0).Format(config.Config.App.TimeLayout))
		if event.Remind {
			resp += fmt.Sprintf(config.Locale.ListEntryRemind,
				time.Unix(event.Time, 0).Sub(time.Unix(event.RemindTime, 0))/time.Minute, event.RemindInterval)
		}
		resp += "\n"

		// each message contains 5 entries
		if i%5 == 4 {
			chanOut <- util.DefaultMsg(qq, resp)
			resp = ""
		}
	}
	if resp != "" {
		chanOut <- util.DefaultMsg(qq, resp)
	}
	if len(events) == 0 {
		chanOut <- util.DefaultMsg(qq, config.Locale.ListEmpty)
	}
	waitGroup.Done()
}

func CheckerRemind(chanOut util.SendMsgChan) {
	for {
		events, err := model.GetEventsToRemind()
		if err != nil {
			log.Println("controller: error when getting remind events")
			log.Panic(err)
		}
		for _, event := range events {
			err = model.DeleteEvent(event.ID.Hex())
			if err != nil {
				log.Println("controller: error when deleting event " + event.ID.Hex())
				log.Panic(err)
			}
			go workerRemind(event, chanOut)
		}
		time.Sleep(30 * time.Second) // check new events every 30s
	}
}

func workerRemind(event model.Event, chanOut util.SendMsgChan) {
	eventTime := time.Unix(event.Time, 0)
	for {
		if time.Now().After(eventTime) {
			break
		}
		chanOut <- util.DefaultMsg(event.User, fmt.Sprintf(config.Locale.Remind, event.Desc,
			time.Unix(event.Time, 0).Sub(time.Now())/time.Minute))
		time.Sleep(time.Duration(event.RemindInterval) * time.Minute)
	}
	chanOut <- util.DefaultMsg(event.User, fmt.Sprintf(config.Locale.RemindExpire, event.Desc))
}
