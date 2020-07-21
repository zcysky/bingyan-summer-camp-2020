/* This file contains functions of updating  */

package controller

import (
	"account/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Update handles request of updating user information
func Update(c *gin.Context) {
	// verify jwt
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" || len(tokenStr) < 7 {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	id, _, valid, err := ParseToken(tokenStr[7:])
	if err != nil {
		failUpdate(c, http.StatusBadRequest, err.Error())
	}
	if !valid {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var newForm model.SignupForm
	err = c.BindJSON(&newForm)
	if err != nil {
		failUpdate(c, http.StatusBadRequest, err.Error())
		return
	}

	err = model.Update(newForm, id)
	if err != nil {
		failUpdate(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"status":  http.StatusCreated,
	})
}

// Delete handles request of deleting user by admin account
func Delete(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" || len(tokenStr) < 7 {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	_, admin, valid, err := ParseToken(tokenStr[7:])
	if err != nil {
		log.Println(err)
	}
	if !valid || !admin {
		failUpdate(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userid := c.Param("userid")

	err = model.Delete(userid)
	if err != nil {
		failUpdate(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted successfully",
		"status":  http.StatusNoContent,
	})
}

func failUpdate(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
		"status":  status,
	})
}
