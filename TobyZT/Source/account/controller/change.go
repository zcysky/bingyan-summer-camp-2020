/* This file contains  */

package controller

import (
	"account/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Update(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
			"status":  401,
		})
		return
	}
	tokenForm, valid, err := ParseToken(tokenStr)
	if err != nil {
		log.Println(err)
	}
	if !valid {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
			"status":  401,
		})
		return
	}
	var updateForm model.UpdateForm
	err = c.BindJSON(&updateForm)
	if err != nil {
		log.Println(err)
	}
	// ----- Not complete yet -----
}

func Delete(c *gin.Context) {

}
