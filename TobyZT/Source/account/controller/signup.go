/* This file contains handler to handle sign-up request */

package controller

import (
	"account/model"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// Signup reads json from post request and check if it's valid
// If valid, sign up a new account in database
func Signup(c *gin.Context) {
	var jsonForm model.SignupJsonForm
	err := c.BindJSON(&jsonForm)
	if err != nil {
		failSignup(c, http.StatusBadRequest, err.Error())
		return
	}
	// verify formats of username, password, email and phone
	form := model.JsonToSignupForm(jsonForm)
	valid, msg := validInfo(form)
	if !valid {
		failSignup(c, http.StatusBadRequest, msg)
		return
	}
	exist, err := model.AccountExist(form.Email)
	if err != nil {
		failSignup(c, http.StatusBadRequest, err.Error())
		return
	}
	if exist {
		failSignup(c, http.StatusForbidden, "User already exists")
		return
	}
	// If token is empty, send a verify email
	if jsonForm.Token == "" {
		hash, err := SHA1Hash(form)
		if err != nil {
			failSignup(c, http.StatusBadRequest, err.Error())
			return
		}
		err = SendEmail(form.Email, "[BingyanSummerCamp] Signup Verification",
			fmt.Sprintf(`%s, you've just made a sign-up request.
			Here's your verification code: %s`, form.Username, hash))
		if err != nil {
			failSignup(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":       "Verification email sent",
			"status":        http.StatusOK,
			"Authorization": "",
		})
		return
	}

	// Verify token in json
	ans, err := SHA1Hash(form)
	if err != nil {
		failSignup(c, http.StatusBadRequest, err.Error())
		return
	}
	if jsonForm.Token != ans {
		failSignup(c,http.StatusUnauthorized,"Wrong token")
		return
	}
	// Sign up in database
	id, err := model.SignupNew(form)

	// Generate jwt
	token, err := GenerateToken(model.TokenForm{
		UserID: id, Email: form.Email, Password: form.Password,
	})

	c.JSON(http.StatusCreated, gin.H{
		"message":       "success",
		"status":        http.StatusCreated,
		"Authorization": "Bearer " + token,
	})
}

// failSignup helps return error info to front end with json
func failSignup(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message":       msg,
		"status":        status,
		"Authorization": "",
	})
}

// validInfo aids to check formats of username, password, email and phone
func validInfo(form model.SignupForm) (valid bool, msg string) {
	// verify email
	ptn := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Email)
	if !valid {
		return false, "Invalid email address"
	}

	// verify phone
	ptn = `^1[3-9]\d{9}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Phone)
	if !valid {
		return false, "Invalid phone number"
	}

	// verify username
	ptn = `^[a-zA-Z0-9_-]{3,12}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Username)
	if !valid {
		return false, "Invalid username"
	}

	// verify password
	ptn = `^[a-z0-9_-]{6,18}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Password)
	if !valid {
		return false, "Invalid password"
	}

	return true, ""
}

// SHA1Hash generates a hash string based on a SignupForm
func SHA1Hash(form model.SignupForm) (hash string, err error) {
	var jwtInfo model.JWTInfo
	err = ParseJson("config/jwt.json", &jwtInfo)

	var dataParams string
	dataParams = form.Email + form.Password + form.Username + form.Phone +
		jwtInfo.Secret
	h := sha1.New()
	h.Write([]byte(dataParams))
	bs := h.Sum(nil)
	hash = hex.EncodeToString(bs)
	return hash, nil
}
