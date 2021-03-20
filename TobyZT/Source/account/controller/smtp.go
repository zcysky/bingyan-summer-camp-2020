package controller

import (
	"account/model"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/mail"
	"net/smtp"
	"os"
)

func SendEmail(dest string, subject string, body string) (err error) {
	// Read configurations from SMTP.json
	f, err := os.Open("config/SMTP.json")
	if err != nil {
		return err
	}
	var contents []byte
	var smtpInfo model.SMTPInfo
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &smtpInfo)
	if err != nil {
		return err
	}

	from := mail.Address{Address: smtpInfo.User}
	to := mail.Address{Address: dest}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	var msg string
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body

	// Establish connection
	server := smtpInfo.Server + ":" + smtpInfo.Port
	host, _, _ := net.SplitHostPort(server)
	auth := smtp.PlainAuth("", smtpInfo.User, smtpInfo.Password, host)
	conn, err := tls.Dial("tcp", server, nil)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()

	ok, _ := c.Extension("AUTH")
	if !ok {
		return fmt.Errorf("AUTH failed")
	}
	err = c.Auth(auth)
	if err != nil {
		return err
	}
	err = c.Mail(from.Address)
	if err != nil {
		return err
	}
	err = c.Rcpt(to.Address)
	if err != nil {
		return err
	}
	w, err := c.Data()
	if w != nil {
		_, err = w.Write([]byte(msg))
	}
	if err != nil {
		return err
	}
	c.Quit()
	return nil
}
