package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mall/model"
	"net/http"
)

//提前拦截，根据Headers中的Authorization内容检查并获取身份信息
func VerifyIdentity(c *gin.Context) {
	log.Println("Checking authorization")

	//获取token内容
	tokenStr := c.Request.Header.Get("Authorization")

	//检查token格式
	if len(tokenStr) < 7 || tokenStr[:6] != "Bearer" {
		fail(c, http.StatusUnauthorized, "token's struct is wrong")
		c.Abort()
		return
	}

	//解析token内容
	userName, err := model.ParseToken(tokenStr[7:])
	if err != nil {
		fail(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	//传递给Handler
	log.Println("Authorization is valid")
	c.Set("username", userName)
	c.Next()
}
