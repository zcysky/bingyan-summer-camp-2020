package controller

import (
	"log"
	"qq-bot-ref/model"
	"sync"
)

type (
	SigChanMap map[string]chan bool
	WaitGroupMap map[string]*sync.WaitGroup
)

func InitController() {
	initEvent()

	err := model.InitRemindingStatus()
	if err != nil {
		log.Println("controller: error when reset reminding status")
		log.Panic(err)
	}
}
