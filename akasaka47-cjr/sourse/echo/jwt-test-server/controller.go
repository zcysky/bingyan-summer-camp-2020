package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	client := Mstart()

	server := echo.New()
	Signup(server,client)


	err := server.Start(":8080")
	if err != nil {
		fmt.Println("server failed")
	}

	Mend(client)
}


