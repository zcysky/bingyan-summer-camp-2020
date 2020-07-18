package module_mapper

import (
	"context"
	"gopkg.in/gomail.v2"
	"time"
	"warmUp/config"
	"warmUp/util"
)

func GenerateRegisterCode(userEmailAddr string) error {

	registerCode, err := util.GenerateUUID()
	if err != nil{
		return err
	}

	RegisterRedis.Set(context.TODO(), registerCode+"_registerCode", registerCode,
		time.Duration(config.Config.Redis.RedisTokenLife) * time.Minute)


	mail := gomail.NewMessage()
	mail.SetHeader("From", config.Config.Mail.MailAddress)
	mail.SetHeader("To", userEmailAddr)
	mail.SetBody("/text/html", "Your register code is <b>" + registerCode + "</b>")

	dialer := gomail.NewDialer(config.Config.Mail.SMTPAddress, config.Config.Mail.MailPort,
		config.Config.Mail.Name, config.Config.Mail.Password)

	if err := dialer.DialAndSend(mail); err != nil{
		return err
	}

	return nil
}

func AuthRegisterCode(code string) (bool, error) {
	val, err := RegisterRedis.Get(context.TODO(), code+"_registerCode").Result()
	if err != nil{
		return false, err
	}

	if val != code {
		return false, nil
	}

	RegisterRedis.Del(context.TODO(), code+"_registerCode")

	return true, nil
}