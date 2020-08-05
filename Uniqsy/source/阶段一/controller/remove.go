package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/model"
)

func Remove(c *gin.Context) {
	// 检查token内容是否正确以及是否为管理员权限
	tokenStr := c.Request.Header.Get("Authorization")
	err := checkAdmin(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":	"failed in token check",
			"error":	err.Error(),
		})
	}

	// 获取个人查询的约束条件即用户在数据库中的唯一id
	var removeForm model.RemoveForm
	err = c.BindJSON(&removeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"query form is wrong",
			"error":	err.Error(),
		})
		return
	}

	err = model.Remove(removeForm.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"failed in removing",
			"error":	err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":	"remove successfully",
		"user_id":	removeForm.UserID,
	})
}