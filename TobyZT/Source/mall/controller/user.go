package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
	"reflect"
	"regexp"
)

// Login handles users' requests of logging in (request with json attached)
func Login(c *gin.Context) {
	var form model.LoginForm
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
	valid = model.VerifyLogin(form)
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

// Signup handles users' requests of creating new accounts (request with json attached)
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

// Update handles users' requests of updating personal info (request with json attached)
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
	valid, errMsg := validInfo(form)
	if !valid {
		failMsg(c, http.StatusUnauthorized, errMsg)
		return
	}

	err = model.UpdateUser(username.(string), form)
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

// LoginVerification is a middleware which ensures users' login status.
// It checks the jwt in the header of request.
func LoginVerification(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	username, valid, err := model.ParseToken(tokenStr[7:])
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

// TryToGetUser is a middleware which checks login status of users.
// It's OK if the user doesn't login. If logged in, set a username context.
// It checks the jwt in the header of request
func TryToGetUser(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	username, _, _ := model.ParseToken(tokenStr[7:])
	if username != "" {
		c.Set("username", username)
	}
	c.Next()
}

func GetSelfInfo(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	form, cnt, err := model.QueryOneUser(username.(string))
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
			"total_view_count":    cnt.ViewCnt,
			"total_collect_count": cnt.CollectCnt,
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
	form, cnt, err := model.QueryOneUser(target)
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
			"total_view_count":    cnt.ViewCnt,
			"total_collect_count": cnt.CollectCnt,
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

func validInfo(form interface{}) (valid bool, errMsg string) {
	t := reflect.TypeOf(form)
	v := reflect.ValueOf(form)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Name {
		case "Email":
			{
				ptn := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
				reg := regexp.MustCompile(ptn)
				valid = reg.MatchString(v.Field(i).String())
				if !valid {
					return false, "invalid email address"
				}
			}
		case "Mobile":
			{
				ptn := `^1[3-9]\d{9}$`
				reg := regexp.MustCompile(ptn)
				valid = reg.MatchString(v.Field(i).String())
				if !valid {
					return false, "invalid phone number"
				}
			}
		case "Username":
			{
				ptn := `^[a-zA-Z0-9_-]{2,12}$`
				reg := regexp.MustCompile(ptn)
				valid = reg.MatchString(v.Field(i).String())
				if !valid {
					return false, "invalid username"
				}
			}
		case "Nickname":
			{
				ptn := `^[a-zA-Z0-9_-]{2,12}$`
				reg := regexp.MustCompile(ptn)
				valid = reg.MatchString(v.Field(i).String())
				if !valid {
					return false, "invalid username"
				}
			}
		case "Password":
			{
				ptn := `^[a-z0-9_-]{6,18}$`
				reg := regexp.MustCompile(ptn)
				valid := reg.MatchString(v.Field(i).String())
				if !valid {
					return false, "invalid password"
				}
			}

		}
	}
	return true, ""
}
