package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/model"
)

//登录
func Login(c *gin.Context) {
	//获取登录信息
	var loginForm model.LoginForm
	err := c.BindJSON(&loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"Wrong form struct",
			"error":	err.Error(),
		})
		return
	}

	//验证登录信息
	id, err := model.VerifyLoginForm(loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"Wrong password or wrong email address"
			"error":	err.Error(),
		})
		return
	}

	//生成对应的token内容
	token, err := GenerateToken(id, loginForm.IsAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":	err.Error(),
		})
		return
	}

	//返回登录结果
	c.JSON(http.StatusOK, gin.H)
}