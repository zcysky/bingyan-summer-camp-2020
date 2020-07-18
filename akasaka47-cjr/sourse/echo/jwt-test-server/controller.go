package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Info Config



func main() {
	InitConfig()
	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(Info.JWTsecret),
	})

	client := Mstart()
	server := echo.New()
	Signup(server, client)
	Login(server, client)
	MainPage(server, IsLoggedIn)

	err := server.Start(":8080")
	if err != nil {
		fmt.Println("server failed")
	}

	Mend(client)
}
