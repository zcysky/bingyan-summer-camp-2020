/* This file contains handlers to handle login request */

package controller

import (
	"account/model"
	"log"

	"github.com/gin-gonic/gin"
)

// Login reads json in the post request and verify it
// If correct, sign a jwt and then return it with json
func Login(c *gin.Context) {
	var form model.LoginForm
	err := c.BindJSON(&form)
	if err != nil {
		log.Println(err)
	}
	valid, admin, id, err := model.VerifyLogin(form)
	if err != nil {
		log.Println(err)
	}
	if valid == false {
		c.JSON(401, gin.H{
			"message":       "failed",
			"status":        401,
			"userid":        id,
			"admin":         0,
			"Authorization": "",
		})
		return
	}

	token, err := GenerateToken(model.TokenForm{
		UserID: id, Email: form.Email, Password: form.Password,
	})
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"message":       "success",
		"status":        200,
		"admin":         admin,
		"Authorization": "bearer " + token,
	})
}
