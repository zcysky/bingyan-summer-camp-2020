package controller

import (
	"net/smtp"
	"strings"
	"task1/config"
)

//重写net/smtp包中的SendEmail函数，因为没法配置SSL
func SendMail(to, subject, body, mailtype string) error {
	auth := smtp.PlainAuth(
		"",
		config.Config.SMTP.User,
		config.Config.SMTP.Password,
		config.Config.SMTP.Server)
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\n" +
		"From: " + config.Config.SMTP.User + "<" + config.Config.SMTP.User + ">\r\n" +
		"Subject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(
		config.Config.SMTP.Server + ":" + config.Config.SMTP.Port,
		auth,
		config.Config.SMTP.User,
		sendTo, msg)
	return err
}