package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var config Config

func main() {
	var err error
	config, err = LoadConfig("./config.json")
	fmt.Println(config)
	if err != nil {
		fmt.Println(err)
	}
	e:=echo.New()
	e.GET("/", GernerateToker)
	e.GET("/check", func(context echo.Context) error {
		id,err:=TokenCheck(context)
		if(err!=nil){
			return context.String(http.StatusBadRequest,"success")
		}
		fmt.Println(id,"<-")//还没返回，暂时直接在终端输出
		return context.String(http.StatusOK,"success")
	})
	e.Start(":1323")
}
