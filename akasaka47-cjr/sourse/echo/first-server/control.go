package main

import (
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	//调用模版引擎渲染模版
	return t.templates.ExecuteTemplate(w, name, data)
}

func hello(server *echo.Echo){
	server.GET("/hw", func(context echo.Context) error {
		fmt.Println("hello")
		return context.String(http.StatusOK, "hello, world!")
	})
}

func random(server *echo.Echo){
	t := &Template{
		//模版引擎支持提前编译模版, 这里对目录下以html结尾的模版文件进行预编译处理
		//预编译处理的目的是为了优化后期渲染模版文件的速度
		templates: template.Must(template.ParseGlob("rand.html")),
	}
	server.Renderer = t
	server.GET("/rand", func(context echo.Context) error {
		seed := context.QueryParam("seed")
		s, err := strconv.ParseInt(seed,10,64)
		if err != nil {
			fmt.Println("failed")
		}
		rand.Seed(s)
		ran := rand.Intn(1000)
		str := strconv.FormatInt(int64(ran),10)
		return context.Render(http.StatusOK, "rand.html", str)
	})
}

func upload(server *echo.Echo)  {
	server.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		str := file.Filename
		fmt.Println("文件名："+str)
		if strings.Contains(str, ".png") || strings.Contains(str, ".jpg"){
			fmt.Println("这是一个图片")
		}else{
			fmt.Println("这不是一个图片")
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
	})
}