/* This file contains query handlers */

package controller

import (
	"account/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAllUsers(c *gin.Context) {
	var form model.QueryForm
	err := c.BindJSON(&form)
	if err != nil {
		failQuery(c, http.StatusBadRequest, err.Error())
		return
	}
	users, err := model.QueryAll(form.Limit, form.Page)
	// ------ Not complete --------
}

func QueryOne(c *gin.Context) {
	userid := c.Param("userid")
	user, err := model.QueryOne(userid)
	if err != nil {
		failQuery(c, http.StatusNotFound, "Not Found such user\n"+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"userInfo": gin.H{
			"userid":   userid,
			"username": user.Username,
			"phone":    user.Phone,
			"email":    user.Email,
		},
	})
}

func failQuery(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
		"status":  status,
	})
}
