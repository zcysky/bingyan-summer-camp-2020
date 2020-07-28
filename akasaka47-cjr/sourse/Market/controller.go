package main

import (
	"fmt"
	"github.com/labstack/echo"
)

var Info Config

func main() {
	InitConfig()

	//var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	//	Claims:     &jwtCustomClaims{},
	//	SigningKey: []byte(Info.JWTsecret),
	//})
	//
	//re, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//}

	client := Mstart()
	server := echo.New()

	User(server, client)
	Commodities(server, client)




	err := server.Start(":8080")
	if err != nil {
		fmt.Println("server failed")
	}

}
