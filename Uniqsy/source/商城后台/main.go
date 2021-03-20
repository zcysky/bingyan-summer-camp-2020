package main

import (
	"mall/config"
	"mall/model"
	"mall/router"
)

func main() {
	config.Init()

	model.InitDataBase()

	router.InitRouter()

	router.RunRouter()
}
