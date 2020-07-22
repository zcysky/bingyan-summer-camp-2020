/* This file contains query handlers */

package controller

import (
	"account/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAllUsers(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" || len(tokenStr) < 7 {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	_, admin, valid, err := ParseToken(tokenStr[7:])
	if err != nil {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if !valid || !admin {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var form model.QueryForm
	err = c.BindJSON(&form)
	if err != nil {
		failQuery(c, http.StatusBadRequest, err.Error())
		return
	}
	users, total, err := model.QueryAll(form.Limit, form.Page)
	if err != nil {
		failQuery(c, http.StatusBadRequest, err.Error())
		return
	}
	var res []gin.H
	for _, user := range users {
		res = append(res, gin.H{
			"userid":   user.UserID,
			"username": user.Username,
			"phone":    user.Phone,
			"email":    user.Email,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"total":   total,
		"page":    form.Page,
		"limit":   form.Limit,
		"users":   res,
	})
}

func QueryOne(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" || len(tokenStr) < 7 {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	_, admin, valid, err := ParseToken(tokenStr[7:])
	if err != nil {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if !valid || !admin {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
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
