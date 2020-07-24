package controller

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"task1/config"
)

//重写net/smtp包中的SendEmail函数，因为没法配置SSL
func sendEmail(destination string, subject string, body string) error {
	//邮件基本信息填写
	from := mail.Address{Address: config.Config.SMTP.User}
	to := mail.Address{Address: destination}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	var msg string
	for key, val := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", key, val)
	}
	msg += "\r\n" + body

	//建立服务器与目标邮箱的连接
	server := config.Config.SMTP.Server + ":" + config.Config.SMTP.Port
	host, _, _ := net.SplitHostPort(server)
	auth := smtp.PlainAuth("",
		config.Config.SMTP.User,
		config.Config.SMTP.Password,
		host)
	connection, err := tls.Dial("tcp", server, nil)
	if err != nil {
		return err
	}
	client, err := smtp.NewClient(connection, host)
	if err != nil {
		return err
	}
	defer client.Close()

	ok, _ := client.Extension("AUTH")
	if !ok {
		return fmt.Errorf("AUTH failed")
	}
	err = client.Auth(auth)
	if err != nil {
		return err
	}
	err = client.Mail(from.Address)
	if err != nil {
		return err
	}
	err = client.Rcpt(to.Address)
	if err != nil {
		return err
	}
	w, err := client.Data()
	if w != nil {
		_, err = w.Write([]byte(msg))
	}
	if err != nil {
		return err
	}
	_ = client.Quit()
	return nil
}