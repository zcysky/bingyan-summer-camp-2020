package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"strconv"
	"warmup/config"
	"warmup/model"
)
const (
	myEmailAddress  = "1426742045@qq.com"
	myEmailPassword = "srmijpjdltokigca"
	SMTPServer      = "smtp.qq.com"
	SMTPServerPort  = 25
)

func SendEmailVerification(email string, verificationCode string) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(myEmailAddress, "ababsad")},
		"To":      {email},
		"Subject": {"Email Verification"},
	})
	m.SetBody("text/plain", verificationCode)
	dialer := gomail.NewDialer(SMTPServer, SMTPServerPort, myEmailAddress, myEmailPassword)
	err := dialer.DialAndSend(m)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func HandleRegister(c echo.Context) error {
	registerInfo := new(config.RegisterInfo)
	err := c.Bind(registerInfo)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取数据")
	}
	newUserUid := c.QueryParam("id")
	newUserCode := c.QueryParam("code")
	if newUserUid != "" {//新用户验证，数据库添加新用户
		newUserExist, err := model.ExistUser(newUserUid)
		if err != nil {
			return c.String(http.StatusInternalServerError, "redis库错误")
		}
		if newUserExist == false {
			return c.String(http.StatusBadRequest, "用户不存在")
		}

		newUserCodeInDataBase,err:=model.FindCode(newUserUid)
		if err != nil {
			return c.String(http.StatusInternalServerError, "redis库错误")
		}
		if(newUserCodeInDataBase!=newUserCode){
			return c.String(http.StatusBadRequest, "邮件验证失败")
		}

		//验证成功，向数据库加入新用户
		registerInfo.Uid = newUserUid
		registerInfo.Type = "general"
		err = model.InsertNewUser(*registerInfo)
		if err != nil {
			return c.String(http.StatusInternalServerError, "无法向数据库添加新用户")
		}

		//在用户成功注册后删除邮件验证码

		err=model.DeleteCode(newUserUid)
		return c.JSON(http.StatusOK, registerInfo)
	} else {
		tmpUid := uuid.NewV4().String()
		verificationCode := rand.Int()
		err:=model.AddCode(tmpUid,verificationCode)
		if(err!=nil){
			return c.String(http.StatusInternalServerError, "redis库错误")
		}
		err = SendEmailVerification(registerInfo.Email, strconv.Itoa(verificationCode))
		if err != nil {
			return c.String(http.StatusInternalServerError, "无法发送验证邮件")
		}
		return c.String(http.StatusOK, "已成功向用户发送邮件，请进行验证,用户id为"+tmpUid)
	}

}
