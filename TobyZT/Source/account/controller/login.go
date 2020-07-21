/* This file contains handlers to handle login request */

package controller

import (
	"account/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login reads json in the post request and verify it
// If correct, sign a jwt and then return it with json
func Login(c *gin.Context) {
	var form model.LoginForm
	err := c.BindJSON(&form)
	if err != nil {
		failLogin(c, http.StatusBadRequest, err.Error())
		return
	}
	valid, admin, id, err := model.VerifyLogin(form)
	if err != nil {
		failLogin(c, http.StatusBadRequest, err.Error())
		return
	}
	if !valid {
		failLogin(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := GenerateToken(model.TokenForm{
		UserID: id, Email: form.Email, Password: form.Password,
	})
	if err != nil {
		failLogin(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"status":        http.StatusOK,
		"admin":         admin,
		"Authorization": "Bearer " + token,
	})
}

// failSignup helps return error info to front end with json
func failLogin(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message":       msg,
		"status":        status,
		"userid":        "",
		"admin":         0,
		"Authorization": "",
	})
}
