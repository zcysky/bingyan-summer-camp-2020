package controller

import (
	"../config"
	"../model"
	"../util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"time"
)

var NotificationTickerJob func(ticker *util.JobTicker)
var NotificationTickerFinishJob func(ticker *util.JobTicker)

func AddNotiToTicker(robot *gomirai.Bot, notiId, eventId int64, userId uint, targetTime int64, interval int64, advance int64) {
	newTicker := util.JobTicker{NotificationTickerFinishJob, NotificationTickerJob,
		notiId, interval, targetTime, userId, eventId,
		advance, robot}
	newTicker.RunAtAdvance()
}

func ResolveTime(text string) (int64, int64, int64, int64, error) {

	//    /addNoti <id> <year>-<month>-<day>-<hour>-<minute> <interval> <advance>
	Regexp := regexp.MustCompile(`^/addNoti\s([\d]+)\s([\d]+)-([\d]+)-([\d]+)-([\d]+)-([\d]+)\s([\d]+)\s([\d]+)$`)
	params := Regexp.FindStringSubmatch(text)
	if len(params) == 9 {
		nowtime := time.Now()
		eventId, _ := strconv.Atoi(params[1])
		year, _ := strconv.Atoi(params[2])
		if year == 0 {
			year = nowtime.Year()
		}
		month, _ := strconv.Atoi(params[3])
		if month == 0 {
			month = int(nowtime.Month())
		}
		day, _ := strconv.Atoi(params[4])
		if day == 0 {
			day = nowtime.Day()
		}
		hour, _ := strconv.Atoi(params[5])
		minute, _ := strconv.Atoi(params[6])
		interval, _ := strconv.Atoi(params[7])
		advance, _ := strconv.Atoi(params[8])
		timeLocation, _ := time.LoadLocation("Asia/Shanghai")
		timeObj := time.Date(year, time.Month(month), day, hour, minute, 0, 0, timeLocation)
		//d, _ := time.ParseDuration("-8h") //UTC与北京时间相差8小时
		//timeObj = timeObj.Add(d)
		timeStamp := timeObj.Unix()
		return timeStamp, int64(interval), int64(advance), int64(eventId), nil
	} else {
		return 0, 0, 0, 0, errors.New("invalid time")
	}
}

//进一步，应该初始化从数据库读取未完成的提醒，定时结束后从数据库删除
func HandleAddNoti(robot *gomirai.Bot, event message.Event, text string) error {
	timeStamp, interval, advance, eventId, err := ResolveTime(text)

	if err != nil {
		_, err_ := robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("can't resolve the parameters\n/addNoti <id> <year>-<month>-<day>-<hour>-<minute> <interval> <advance>"))
		if err_ != nil {
			return err_
		}
		return err
	}
	timeLayout := "2006-01-02 15:04:05"
	fmt.Println(time.Unix(timeStamp, 0).Format(timeLayout), time.Now())
	fmt.Println(timeStamp, time.Now().Unix())
	if timeStamp < time.Now().Unix() { //时间检查
		_, err := robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("invalid time"))
		if err != nil {
			return err
		}
		return nil
	}

	//fmt.Println(params,time.Now().Hour(),time.Now().Minute(),hour)
	tmpNotiId := config.ConfigSet.EventCountConfig.Id
	newNoti := model.Notification{event.Sender.Id, eventId,
		timeStamp, interval, advance, tmpNotiId}
	config.ConfigSet.EventCountConfig.Id++
	err = model.AddNoti(newNoti, fmt.Sprint(event.Sender.Id))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("there is no such event"))
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	//fmt.Println("add notification to ticker successfully")
	_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("add Notification successfully"))
	if err != nil {
		return err
	}
	AddNotiToTicker(robot, tmpNotiId, int64(eventId), event.Sender.Id, timeStamp, int64(interval), int64(advance))
	return nil
}

func SendNotification(robot *gomirai.Bot, userId uint, eventId int64) {
	_, err := robot.SendFriendMessage(userId, 0, message.PlainMessage("may i have your notice?"))
	fmt.Println(userId, eventId, fmt.Sprint(userId), "->>>", strconv.FormatInt(eventId, 10))
	eventInfo, err := model.Find(fmt.Sprint(userId), strconv.FormatInt(eventId, 10))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = robot.SendFriendMessage(userId, 0, message.PlainMessage("can't find the document that should be informed"))
			if err != nil {
				fmt.Println(err)
			}
			return

		} else {
			fmt.Println(err)
			return
		}
	}
	eventInfoJSON, err := json.Marshal(eventInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = robot.SendFriendMessage(userId, 0, message.PlainMessage(string(eventInfoJSON)))
	if err != nil {
		return
	}
}

func LoadNotification() error {
	robot, err := ApplyRobot()
	if err != nil {
		return err
	}
	allCol, err := model.ShowAllCollection()
	if err != nil {
		return err
	}
	for _, Col := range allCol {
		allNoti, err := model.ShowAllNoti(Col)
		if err != nil {
			return err
		}
		for _, noti := range allNoti {
			timeStamp := time.Now().Unix()
			if noti.Time < timeStamp {
				err := model.DeleteNoti(noti.NotiId, fmt.Sprint(noti.UserId))
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			fmt.Println("load notification", Col, noti)
			AddNotiToTicker(robot, noti.NotiId, noti.EventId, noti.UserId, noti.Time, noti.Interval, noti.Advance)
		}
	}
	fmt.Println("load notification successfully")
	return nil
}

func ApplyRobot() (*gomirai.Bot, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := gomirai.NewClient("default", config.ConfigSet.MiraiConfig.ClientHost, config.ConfigSet.MiraiConfig.AuthKey)
	c.Logger.Level = logrus.TraceLevel
	key, err := c.Auth()
	if err != nil {
		return nil, err
	}
	b, err := c.Verify(config.ConfigSet.MiraiConfig.QQNumber, key)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func init() {
	NotificationTickerJob = func(ticker *util.JobTicker) {
		SendNotification(ticker.Robot, ticker.TargetId, ticker.TargetEvent)
		//fmt.Println("ticker perform well")
	}
	NotificationTickerFinishJob = func(ticker *util.JobTicker) {
		err := model.DeleteNoti(ticker.NotiId, fmt.Sprint(ticker.TargetId))
		//fmt.Println(ticker.NotiId, ticker.TargetId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				_, err = ticker.Robot.SendFriendMessage(ticker.TargetId, 0, message.PlainMessage("can't find the notification that should be deleted"))
			}
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	err := LoadNotification()
	if err != nil {
		fmt.Println(err)
	}
}
