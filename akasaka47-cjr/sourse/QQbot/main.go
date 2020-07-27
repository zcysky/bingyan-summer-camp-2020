package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
)

var M MessageModel
var Info Config
var RemindID string
var T RespondType

func main() {

	InitConfig()
	InitType()
	client := Mstart()

	re, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}

	server := echo.New()
	go test(server, client, re)

	err = server.Start(":5700")
	if err != nil {
		fmt.Println("server failed")
	}

}
