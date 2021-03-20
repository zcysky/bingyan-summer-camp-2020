package controller

import (
	"../config"
	"encoding/json"
	"fmt"
	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)
import "../model"

func HandleAddEvent(userId string, eventInfo message.Event) error {
	err := model.InsertMemo(userId, eventInfo)
	if err != nil {
		return err
	}
	return nil
}

func HandleQuery(userId string) ([]message.Event, error) {
	allEvent, err := model.ShowAllEvent(userId)
	if err != nil {
		return []message.Event{}, err
	}
	return allEvent, nil
}

func HandleFind(userId string, eventId string) (message.Event, error) {
	Event, err := model.Find(userId, eventId)
	if err != nil {
		return message.Event{}, err
	}
	return Event, nil
}

func HandleDeleteEvent(userId string, eventId string) error {
	err := model.Delete(userId, eventId)
	if err != nil {
		return err
	}
	return nil
}

func HandleSave(robot *gomirai.Bot, event message.Event) error {

	//fmt.Println(e.Sender.Id, fmt.Sprint(e.Sender.Id))
	tmpE := event
	tmpE.EventId = uint(config.ConfigSet.EventCountConfig.Id)
	config.ConfigSet.EventCountConfig.Id++
	err := HandleAddEvent(fmt.Sprint(event.Sender.Id), tmpE)
	if err != nil {
		return err
	}
	_, err = robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("your event has been saved successfully"))
	if err != nil {
		return err
	}
	return nil
}

func HandleShowAll(robot *gomirai.Bot, event message.Event) error {
	_, err := robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("your event is shown in the following text"))
	if err != nil {
		return err
	}
	allEvent, err := HandleQuery(fmt.Sprint(event.Sender.Id))
	if err != nil {
		fmt.Println(err)
	}
	for _, allEventOne := range allEvent {
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


func HandleDelete(robot *gomirai.Bot, event message.Event, text string) error {

	Regexp := regexp.MustCompile(`^/delete\s([\d]+)$`)
	params := Regexp.FindStringSubmatch(text)
	//fmt.Println(params)
	if len(params) == 2 {
		err := HandleDeleteEvent(fmt.Sprint(event.Sender.Id), params[1])
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
	}else {
		_, err := robot.SendFriendMessage(event.Sender.Id, 0, message.PlainMessage("can't resolve the parameters"))
		if err != nil {
			return err
		}

	}
	return nil
}
