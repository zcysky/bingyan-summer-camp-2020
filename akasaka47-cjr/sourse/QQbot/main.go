package main

import (
	"fmt"
	"github.com/labstack/echo"
)

var M MessageModel
var Info Config
var RemindID string

func main() {

	InitConfig()
	client := Mstart()

	server := echo.New()
	test(server, client)

	err := server.Start(":5700")
	if err != nil {
		fmt.Println("server failed")
	}

}
