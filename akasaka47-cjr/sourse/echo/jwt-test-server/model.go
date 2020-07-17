package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
	"strings"
)

func Signup(server *echo.Echo, client *mongo.Client) {
	re, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	server.POST("/emailVerification", func(context echo.Context) error {
		//test()
		u := new(User)
		if err := context.Bind(u); err != nil {
			fmt.Println("error")
		}
		//Insert(client,u)
		mail := new(Email)
		mail.Name = u.Email
		if strings.Count(mail.Name, "@") == 1 {
			if Check(client, u) {
				mail.Status = true
				mail.Info = "该邮箱已被注册！"
			} else {
				mail.Status = false
				mail.Info = "该邮箱未被注册！验证码已发送，请查收。"
				str := CreateRandomString(6)
				//emailReg := utils.Email{
				//	Username: "akasaka907@163.com",
				//	Password: "ryuunosuke47",       //"VWBWKCXALLJBRGIR",
				//	Host: "smtp.163.com",
				//	Port: 25,
				//}
				//emailReg.Subject = "echo服务器测试，注册激活码" //标题：某某软件激活
				//emailReg.From = str
				//emailReg.To = []string{mail.Name}
				//err := emailReg.Send()
				//if err != nil {
				//	mail.Checkstr = "send email failed"
				//}
				err := SendMail(mail.Name, "test-echo验证码", str)
				if err != nil {
					mail.Checkstr = "send email failed"
				} else {
					mail.Checkstr = "send email success"
					err := SetRedis(re, mail.Name, str, "60")
					if err != nil {
						fmt.Println("set-redis failed")
					}
				}
			}
		} else {
			mail.Status = false
			mail.Info = "邮箱输入错误，请输入正确的邮箱名"
		}
		return context.JSON(http.StatusOK, mail)
	})
	server.POST("/signup", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			fmt.Println("error")
		}
		status, err:=FindRedis(re, u.Email, u.Checkstr)
		if err == nil || status == true{
			u.ID = CreateRandomString(4)
			u.Checkstr = ""
			err := Insert(client,u)
			if err != nil {
				u.Info = "用户创建失败"
			}else{
				u.Info = "用户创建成功"
			}
		}else{
			u.Info = "验证码错误，用户创建失败"
		}
		u.Password = ""
		return context.JSON(http.StatusOK, u)
	})
}

func Login(server *echo.Echo) {

}

func SendMail(mailTo string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码:sltdozgngjeacigf
	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}
	mailConn := map[string]string{
		"user": "3250237515@qq.com",
		"pass": "sltdozgngjeacigf",
		"host": "smtp.qq.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "echo-test-mail"))
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo)       //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/plain", body)   //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
