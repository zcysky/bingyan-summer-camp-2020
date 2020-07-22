package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

func test(server *echo.Echo, client *mongo.Client) {
	server.POST("/", func(context echo.Context) error {
		m := new(PostInfo)
		if err := context.Bind(m); err != nil {
			return err
		}
		//fmt.Println(m)
		fmt.Println(m.Message)
		fmt.Println(M.Tag)
		r := new(Respond)
		if strings.Contains(m.Message, "") {
			Switch(m, r)
			YYGQ(m, r)
			Reminder(m, r, client)
		}
		//fmt.Println(r)
		return context.JSON(http.StatusOK, r)
	})
}
