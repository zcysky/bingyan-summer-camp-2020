package main

import (
	"fmt"
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
	//a bit of different from previous
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	fmt.Println("fine")
	match, err := ImgFormat(avatar.Filename)
	if err != nil {
		return c.String(http.StatusBadRequest, "匹配问题")
	}
	if !match {
		return c.String(http.StatusBadRequest, "不是图片格式")
	}
	fmt.Println("fine")
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("./" + avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "图片已保存")

}

func main() {
	e := echo.New()
	e.GET("/hw", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/rand", func(c echo.Context) error {
		seed, _ := strconv.Atoi(c.QueryParam("seed"))
		rand.Seed(int64(seed))
		randNum := rand.Int()
		return c.String(http.StatusOK, strconv.Itoa(randNum))
	})
	e.POST("/image", PostImage)
	e.Start(":1323")
}
