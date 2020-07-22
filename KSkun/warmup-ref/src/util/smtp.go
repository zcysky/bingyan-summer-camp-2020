package util

import (
	"github.com/go-gomail/gomail"
	"log"
	"os"
	"strconv"
)

var (
	smtpDialer   *gomail.Dialer
	emailAddress string
)

func initUtilSMTP() {
	var found bool
	emailAddress, found = os.LookupEnv("EMAIL_ADDR")
	emailPassword, found := os.LookupEnv("EMAIL_PASS")
	smtpAddress, found := os.LookupEnv("SMTP_ADDR")
	smtpPortStr, found := os.LookupEnv("SMTP_PORT")
	if !found {
		log.Println("util: smtp env variables not found")
		log.Panic()
	}

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Println("util: cannot parse smtp port " + smtpPortStr)
		log.Panic(err)
	}

	smtpDialer = gomail.NewDialer(smtpAddress, smtpPort, emailAddress, emailPassword)
}

func SendEmail(receiver string, subject string, content string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", emailAddress)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", content)
	return smtpDialer.DialAndSend(mail)
}
