package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//这个能匹配 /user/tom , 但是不能匹配 /user/ 或  /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//有一个方法可以匹配 /user/tom, 也可以匹配 /user/tom/send
	//如果没有任何了路由匹配 /user/tom, 它将会跳转到 /user/tom/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.Run(":8080")
}