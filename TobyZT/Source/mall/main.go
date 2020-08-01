package main

import (
	"log"
	"mall/model"
	"mall/router"
)

func main() {

	err := model.SetupDatabase()
	if err != nil {
		log.Println(err)
	}
	r := router.InitRouter()
	r.Run(":3939")

}
