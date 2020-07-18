/* This file contains handler to handle sign-up request */

package controller

import (
	"account/model"
	"github.com/gin-gonic/gin"
	"log"
)

// Signup reads json from post request and check if it's valid
// If valid, sign up a new account in database
func Signup(c *gin.Context) {
	var form model.SignupForm
	err := c.BindJSON(&form)
	if err != nil {
		log.Println(err)
	}
	valid := validInfo(form)
	if valid == false {
		c.JSON(400, gin.H{
			"message":       "Invalid user information",
			"status":        400,
			"Authorization": "",
		})
		return
	}
	exist, err := model.SignupNew(form)
	if err != nil {
		log.Println(err)
	}
	if exist == true {
		c.JSON(403, gin.H{
			"message":       "User already exists",
			"status":        403,
			"Authorization": "",
		})
		return
	}
	token, err := GenerateToken(model.LoginForm{
		Email: form.Email, Password: form.Password,
	})
	c.JSON(201, gin.H{
		"message":       "success",
		"status":        201,
		"Authorization": "bearer " + token,
	})
}

// validInfo aids to check whether the info in signup form is valid
func validInfo(form model.SignupForm) (valid bool) {

	return true
}
