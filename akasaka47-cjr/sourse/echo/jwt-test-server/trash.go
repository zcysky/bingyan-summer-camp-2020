package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//
//import (
//	_ "fmt"
//	"github.com/labstack/echo"
//	"html/template"
//	"io"
//	_ "math/rand"
//	"net/http"
//	_ "strconv"
//)
//
//type Template struct {
//	templates *template.Template
//}
//
//func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//	//调用模版引擎渲染模版
//	return t.templates.ExecuteTemplate(w, name, data)
//}
//
//func tinits(server *echo.Echo){
//	t := &Template{
//		templates: template.Must(template.ParseGlob("*.html")),
//	}
//	server.Renderer = t
//	server.GET("/init", func(context echo.Context) error {
//		return context.Render(http.StatusOK, "init.html", "")
//	})
//}
//
//func tcode(server *echo.Echo){
//	server.GET("/jwt", func(context echo.Context) error {
//		username := context.QueryParam("username")
//		//password := context.QueryParam("password")
//		user := Users{Name: username,
//		}
//		str := NewTokens(username, random())
//		decode(str)
//		cookie := new(http.Cookie)
//		cookie.Name = "JWT"
//		cookie.Value = str
//		cookie.Path = "/jwt"
//		//cookie有效期为3600秒
//		cookie.MaxAge = 3600
//		//设置cookie
//		context.SetCookie(cookie)
//
//		return context.Render(http.StatusOK, "ok.html", map[string]interface{}{
//			"user": user,
//			"str": str,
//		})
//	})
//}

func test(){
	c, err := redis.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}
