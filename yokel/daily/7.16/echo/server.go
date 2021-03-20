package main

import (
	"github.com/labstack/echo"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	rgex = `*\.jpg|*\.pg|*\.gif|*\.jpeg$`
)

func ImgFormat(img string) (bool, error) {

	match, err := regexp.MatchString("png", img)
	if err != nil {
		return false, err
	}
	return match, nil
}

func PostImage(c echo.Context) error {
	//a bit of different from previous function,when using api of echo
	//we are already on the top of the program
	//we need to fix the erro here rather than throw up

	//open file in POST body in form-data format
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return c.String(http.StatusInternalServerError,"无法读取图片")
	}
	//fmt.Println("fine")
	match, err := ImgFormat(avatar.Filename)
	if err != nil {
		return c.String(http.StatusBadRequest, "匹配问题")
	}
	if !match {
		return c.String(http.StatusBadRequest, "不是图片格式")
	}
	//fmt.Println("fine")
	src, err := avatar.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError,"无法打开图片")
	}
	defer src.Close()

	dst, err := os.Create("./" + avatar.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError,"无法创建文件")
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return c.String(http.StatusInternalServerError,"无法保存图片")
	}
	return c.HTML(http.StatusOK, "图片已保存")

}

func main() {
	e := echo.New()
	e.GET("/hw", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/rand", func(c echo.Context) error {
		seed, err := strconv.Atoi(c.QueryParam("seed"))
		if err!=nil {
			return c.String(http.StatusBadRequest,"无法解析种子")
		}
		rand.Seed(int64(seed))
		randNum := rand.Int()
		return c.String(http.StatusOK, strconv.Itoa(randNum))
	})
	e.POST("/image", PostImage)
	e.Start(":1323")
}


//some notes of GET and POST may help

//GET is used to request data from a specified resource.
//GET requests can be cached
//GET requests remain in the browser history
//GET requests can be bookmarked
//GET requests should never be used when dealing with sensitive data
//GET requests have length restrictions
//GET requests are only used to request data (not modify)

//POST is used to send data to a server to create/update a resource.
//POST requests are never cached
//POST requests do not remain in the browser history
//POST requests cannot be bookmarked
//POST requests have no restrictions on data length
