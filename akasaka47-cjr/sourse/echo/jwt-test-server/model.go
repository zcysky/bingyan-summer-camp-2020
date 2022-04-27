package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gomail.v2"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
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
			if CheckEmail(client, u) {
				mail.Status = false
				mail.Info = "该邮箱已被注册！"
			} else {
				mail.Status = true
				mail.Info = "该邮箱未被注册！验证码已发送，请查收。"
				str := CreateRandomString(6)
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
		context.Handler()
		status, err := FindRedis(re, u.Email, u.Checkstr)
		if CheckEmail(client, u) {
			u.Info = "用户已存在"
			u.Status = false
		} else {
			if err == nil && status == true {
				u.ID = CreateRandomString(4)
				u.Checkstr = ""
				u.Info = "用户创建成功"
				u.Status = true
				err := Insert(client, u)
				if err != nil {
					u.Info = "用户创建失败"
				}
			} else {
				u.Status = false
				u.Info = "验证码错误，用户创建失败"
			}
		}
		u.Password = ""
		return context.JSON(http.StatusOK, u)
	})
}

func Login(server *echo.Echo, client *mongo.Client) {
	server.POST("/login", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			fmt.Println("error")
		}
		status, result := CheckUser(client, u)
		if status {
			// Set custom claims
			claims := &jwtCustomClaims{
				result.Name,
				true,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
					Id: result.ID,
				},
			}
			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Generate encoded token and send it as response.
			tokenstr, err := token.SignedString([]byte(Info.JWTsecret))
			if err != nil {
				return err
			}
			result.Password = ""
			result.Info = "登录成功"
			result.Status = true
			result.Token = tokenstr
			return c.JSON(http.StatusOK, result)
		}
		ru := new(User)
		ru.Status = false
		ru.Info = "登录失败"
		return c.JSON(http.StatusOK, ru)
	})
}

func MainPage(server *echo.Echo, IsLoggedIn echo.MiddlewareFunc) {
	server.GET("/mainpage", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name
		r := new(User)
		r.Status = true
		r.Info = "登录成功"
		r.Name = name
		r.ID = claims.Id
		return c.JSON(http.StatusOK, r)
	}, IsLoggedIn)
}

func CreateRandomString(len int) string {
	var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var str = "1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func SetRedis(re redis.Conn, mail string, check string, time string) error {
	_, err := re.Do("SET", "mail", mail, "EX", time)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	_, err = re.Do("SET", "check", check, "EX", time)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	fmt.Println("set-redis success")
	return nil
}

func FindRedis(re redis.Conn, mail string, check string) (bool, error) {
	m, err := redis.String(re.Do("GET", "mail"))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return false, err
	}
	c, err := redis.String(re.Do("GET", "check"))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return false, err
	}
	if mail == m && c == check {
		return true, nil
	}
	return false, nil
}

func SendMail(mailTo string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}
	mailConn := map[string]string{
		"user": Info.MailName,
		"pass": Info.MainAuth,
		"host": Info.MailHost,
		"port": Info.MailPort,
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