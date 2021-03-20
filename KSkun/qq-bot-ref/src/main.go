package main

import (
	"log"
	"qq-bot-ref/config"
	"qq-bot-ref/controller"
	"qq-bot-ref/model"
	"qq-bot-ref/util"
	"qq-bot-ref/view"
)

func main() {
	log.Println("qq-bot-ref " + config.VERSION + " starting...")

	util.InitUtil()
	config.InitConfig()
	config.InitLocale()
	model.InitModel()
	controller.InitController()
	view.InitView()
}
