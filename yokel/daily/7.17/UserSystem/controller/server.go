package controller

import (
	"../config"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/user", HandleRegister)
	fmt.Println(config.Config.JWT.Secret)

	//userJwt := e.Group("/user", middleware.JWT(config.Config.JWT.Secret))
	//
	//userJwt.GET("", func(context echo.Context) error {
	//	//do something
	//	return nil
	//})

	e.Start(":1323")
}
