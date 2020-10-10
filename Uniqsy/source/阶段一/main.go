package main

import (
	"task1/config"
	"task1/model"
	"task1/router"
)

func main() {
	config.Init()

	model.SetupDatabase()

	r := router.InitRouter()

	r.Run(":8080")
}
