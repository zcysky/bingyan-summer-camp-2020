package main

import (
	"./config"
	"./controller"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.POST("/register-form", controller.HandleRegister)

	fmt.Println(config.Config.JWT.Secret)

	userJwt := e.Group("/user", middleware.JWT(config.Config.JWT.Secret))

	userJwt.GET("", func(context echo.Context) error {
		//do something
		return nil
	})

	e.Start(":1323")
}
