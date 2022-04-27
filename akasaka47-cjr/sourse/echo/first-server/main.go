package main

import (
	//导入echo包
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	//实例化echo对象。
	server := echo.New()
	hello(server)
	random(server)
	upload(server)

	err := server.Start(":8080")
	if err != nil {
		fmt.Println("server failed")
	}
}