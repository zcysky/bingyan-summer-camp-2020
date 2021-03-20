package controller

import "github.com/Logiase/gomirai/message"
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
		return []message.Event{},err
	}
	return allEvent, nil
}

func HandleFind(userId string,eventId string) (message.Event, error) {
	Event, err := model.Find(userId,eventId)
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
