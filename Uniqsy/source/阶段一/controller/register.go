package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"task1/config"
	"task1/model"
)

func Register(c *gin.Context) {
	//读取前端获取的粗信息
	var rawForm model.RawRegisterForm
	err := c.BindJSON(&rawForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":	err.Error(),
		})
	}

	//将获取的信息精简为在数据库中存储的形式
	dbForm := model.GetDBRegisterForm(rawForm)

	//检查该邮箱是否已经被注册
	exist := model.CheckEmail(dbForm.Email)
	if exist == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"Email address is invalid已注册",
		})
		return
	}

	//给邮箱发一封邮件进行SMTP验证
	subject := "用户管理系统注册"
	body := "欢迎注册，这是一封验证邮件"
	err = SendMail(dbForm.Email, subject, body, "text")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"Email address is invalid",
		})
		return
	}

	//对于要注册管理员的请求，比对邀请码
	//将信息存入，然后获取存入的id
	var collection string
	if rawForm.IsAdmin {
		if strings.Compare(rawForm.Invitation, config.Config.Invitation) == 0{
			collection = "admin"
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":	"invitation code wrong",
			})
			return
		}
	} else {
		collection = "users"
	}
	id, err := model.SaveInDB(dbForm, collection)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":	"Registered successfully",
		"id":		id,
	})
}