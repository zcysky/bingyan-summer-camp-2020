package main

import (
	"task1/model"
	"task1/router"
)

func main() {
	model.SetupDatabase()

	r := router.InitRouter()

	r.Run(":8080")
}
