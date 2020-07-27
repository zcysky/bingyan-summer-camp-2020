package util

import (
	"errors"
	"fmt"
	"qq-bot-ref/config"
)

type MsgToSend struct {
	Msg string
	QQ  uint
}

type SendMsgChan chan MsgToSend

func DefaultMsg(qq uint, msg string) MsgToSend {
	return MsgToSend{
		Msg: msg,
		QQ:  qq,
	}
}

func SuccessMsg(qq uint, funcName string) MsgToSend {
	return DefaultMsg(qq, fmt.Sprintf(config.Locale.Success, funcName))
}

func FailedMsg(qq uint, funcName, msg string) MsgToSend {
	if config.Config.App.Debug {
		return DefaultMsg(qq, fmt.Sprintf(config.Locale.FailedDebug, funcName, msg))
	}
	return DefaultMsg(qq, fmt.Sprintf(config.Locale.Failed, funcName))
}

func Boolean(msg string) (bool, error) {
	if msg == "是" || msg == "Y" || msg == "y" {
		return true, nil
	}
	if msg == "否" || msg == "N" || msg == "n" {
		return false, nil
	}
	return false, errors.New("util: " + msg + " is not a boolean type message")
}
