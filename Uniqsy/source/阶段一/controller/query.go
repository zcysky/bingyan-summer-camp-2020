package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/model"
)

func QueryAll(c *gin.Context) {
	// 检查token内容是否正确以及是否为管理员权限
	tokenStr := c.Request.Header.Get("Authorization")
	err := checkAdmin(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":	"failed in token check",
			"error":	err.Error(),
		})
	}

	// 获取查询的约束条件并检查
	var queryForm model.QueryForm
	err = c.BindJSON(&queryForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"query form is wrong",
			"error":	err.Error(),
		})
		return
	}

	// 进行查询
	users, total, err := model.QueryAll(queryForm.Limit, queryForm.Page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"failed in querying",
			"error":	err.Error(),
		})
		return
	}

	//查询结果写入
	var res []gin.H
	for _, user := range users {
		res = append(res, gin.H{
			"user_id":   user.UserID,
			"user_name": user.Username,
			"tel":    user.Phone,
			"email":    user.Email,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result":	"Query all users successfully",
		"total":   	total,
		"limit":   	queryForm.Limit,
		"page":   	queryForm.Page,
		"users":  	res,
	})
}

func QueryOne(c *gin.Context) {
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
	var queryForm model.QueryForm
	err = c.BindJSON(&queryForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"query form is wrong",
			"error":	err.Error(),
		})
		return
	}

	user, err := model.QueryOne(queryForm.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":	"failed in querying",
			"error":	err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":	"query successfully",
		"user":	gin.H{
			"user_id":		user.UserID,
			"user_name":	user.Username,
			"user_phone":	user.Phone,
			"user_email":	user.Email,
		},
	})
}

func checkAdmin(tokenStr string) (err error) {
	isAdmin, err := ParseToken(tokenStr)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errors.New("you are not admin user")
	}
	return nil
}