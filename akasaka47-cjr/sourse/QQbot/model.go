package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

func test(server *echo.Echo, client *mongo.Client, re redis.Conn) {
	server.POST("/", func(context echo.Context) error {
		m := new(PostInfo)
		if err := context.Bind(m); err != nil {
			return err
		}
		//fmt.Println("+++", m, "+++")
		//fmt.Println(m.GroupId)
		//fmt.Println(M.Tag)
		r := new(Respond)
		if strings.Contains(m.Message, "") {
			YYGQ(m, r)
			Reminder(m, r, client, re)
			//SetRemindTime(m, r, client, re)
			Switch(m, r)
		}
		//fmt.Println(r)
		return context.JSON(http.StatusOK, r)
	})
}

func SetRedis(re redis.Conn, key string, value string) error {
	_, err := re.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	fmt.Println("set-redis success")
	return nil
}

func DelRedis(re redis.Conn, key string){
	_, err := re.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delete failed:", err)
	}
}

func FindRedis(re redis.Conn, key string) string{
	result, err := redis.String(re.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return ""
	}
	return result
}