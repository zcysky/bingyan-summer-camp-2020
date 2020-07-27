package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
	"regexp"
)

func Login(c *gin.Context) {
	var form model.LoginForm
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	valid := model.VerifyLogin(form)
	if !valid {
		failMsg(c, http.StatusUnauthorized, "wrong username or password")
		return
	}
	token, err := model.GenerateToken(form.Username)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data":    "Bearer " + token,
	})

}

func Signup(c *gin.Context) {
	var form model.User
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	valid, errMsg := validInfo(form)
	if !valid {
		failMsg(c, http.StatusBadRequest, errMsg)
		return
	}
	err = model.Signup(form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"error":   "",
		"data":    "",
	})
}

func Update(c *gin.Context) {
	var form model.UpdateForm
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	valid, errMsg := validInfo(model.User{Username: "default", Password: form.Password,
		Nickname: form.Nickname, Mobile: form.Mobile, Email: form.Email})
	if !valid {
		failMsg(c, http.StatusUnauthorized, errMsg)
		return
	}

	err = model.Update(username.(string), form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok",
	})
}

func LoginVerification(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	username, valid, err := model.ParseToken(tokenStr)
	if !valid || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	c.Set("username", username)
	c.Next()
}

func GetSelfInfo(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	form, err := model.QueryOne(username.(string))
	if err != nil {
		failMsg(c, http.StatusGone, "user not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"username":            form.Username,
			"nickname":            form.Nickname,
			"mobile":              form.Mobile,
			"email":               form.Email,
			"total_view_count":    0,
			"total_collect_count": 0,
		},
	})
}

func GetPublicInfo(c *gin.Context) {
	self, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	target := c.Param("id")
	if target == self {
		GetSelfInfo(c)
		return
	}
	form, err := model.QueryOne(target)
	if err != nil {
		failMsg(c, http.StatusGone, "user not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"nickname":            form.Nickname,
			"email":               form.Email,
			"total_view_count":    0,
			"total_collect_count": 0,
		},
	})
}

func failMsg(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   msg,
		"data":    "",
	})
	c.Abort()
}

func validInfo(form model.User) (valid bool, errMsg string) {
	// verify email
	ptn := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Email)
	if !valid {
		return false, "invalid email address"
	}

	// verify phone
	ptn = `^1[3-9]\d{9}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Mobile)
	if !valid {
		return false, "invalid phone number"
	}

	// verify username
	ptn = `^[a-zA-Z0-9_-]{3,12}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Username)
	if !valid {
		return false, "invalid username"
	}
	valid = reg.MatchString(form.Nickname)
	if !valid {
		return false, "invalid username"
	}

	// verify password
	ptn = `^[a-z0-9_-]{6,18}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Password)
	if !valid {
		return false, "invalid password"
	}

	return true, ""
}
