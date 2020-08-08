package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
	"regexp"
)

//注册用户
func Register(c *gin.Context) {
	//获取前端信息并解析到struct
	var registerForm model.RegisterForm
	err := c.BindJSON(&registerForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查数据有效性
	err = checkRegisterInfo(registerForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查用户名、邮箱是否已经注册
	err = model.CheckExist(registerForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//将注册信息加入数据库
	err = model.Register(registerForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	successStr(c, http.StatusCreated, "")
}

//用户登录，并获取token
func Login(c *gin.Context) {
	//获取前端信息并解析到struct
	var loginForm model.LoginForm
	err := c.BindJSON(&loginForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查数据有效性
	err = checkLoginInfo(loginForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查数据库中是否有该账户
	err = model.VerifyLogin(loginForm)
	if err != nil {
		fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	//生成加密后的字符串token
	tokenStr, err := model.GenerateToken(loginForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	successStr(c, http.StatusOK, "Bearer" + tokenStr)
}

//根据用户名获取用户信息
func GetOthersInfo(c *gin.Context) {
	//获取经过中间件处理的username
	selfUserName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//获取URL中给出的用户名，如果是本人，则转到GetMeInfo接口
	targetUserName := c.Param("id")
	if selfUserName == targetUserName {
		GetMeInfo(c)
		return
	}

	//查询用户信息
	targetUser, targetCnt, err := model.QueryUser(targetUserName)
	if err != nil {
		fail(c, http.StatusNotFound, err.Error())
		return
	}

	//提取有用的信息并返回
	data := gin.H{
		"nickname":	targetUser.NickName,
		"email":	targetUser.Email,
		"total_view_count":	targetCnt.ViewCnt,
		"total_collect_count":	targetCnt.CollectCnt,
	}
	successH(c, http.StatusOK, data)
}

//正则表达式检查注册表单的内容格式
func checkRegisterInfo(form model.RegisterForm) error {
	//检查邮箱格式
	ptn := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(ptn)
	valid := reg.MatchString(form.Email)
	if !valid {
		return errors.New("Invalid email address")
	}

	//检查手机号格式
	ptn = `^1[3-9]\d{9}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Mobile)
	if !valid {
		return errors.New("invalid phone number")
	}

	//检查用户名格式
	ptn = `^[a-zA-Z0-9_-]{2,12}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.UserName)
	if !valid {
		return errors.New("invalid username")
	}

	//检查昵称格式
	valid = reg.MatchString(form.NickName)
	if !valid {
		return errors.New("invalid username")
	}

	//检查密码格式
	ptn = `^[a-z0-9_-]{6,18}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(form.Password)
	if !valid {
		return errors.New("invalid password")
	}

	return nil
}

//正则表达式检查登录表单的内容格式
func checkLoginInfo(loginForm model.LoginForm) (err error) {
	//检查用户名格式
	ptn := `^[a-zA-Z0-9_-]{2,12}$`
	reg := regexp.MustCompile(ptn)
	valid := reg.MatchString(loginForm.UserName)
	if !valid {
		return errors.New("invalid username")
	}

	//检查密码格式
	ptn = `^[a-z0-9_-]{6,18}$`
	reg = regexp.MustCompile(ptn)
	valid = reg.MatchString(loginForm.Password)
	if !valid {
		return errors.New("invalid password")
	}

	return nil
}