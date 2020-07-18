package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/hw", func(c *gin.Context) {
		//返回HelloWorld内容
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/rand/:seed", func(c *gin.Context) {
		//根据提供的seed随机生成大数
		seedStr := c.Param("seed")
		seed, err := strconv.ParseInt(seedStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		rand.Seed(seed)
		randNum := rand.Int63()
		c.String(http.StatusOK, strconv.FormatInt(randNum, 10))
	})

	r.POST("/image", func(c *gin.Context) {
		//上传图片并检查图片类型

		//获取图片名
		fileName := c.PostForm("name")
		//获取图片文件
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		//打开图片，
		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		//读取图片的前512个字节，用于鉴定文件是否为图片
		buffer := make([]byte, 512)
		_, err = f.Read(buffer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		f.Close()

		//鉴定文件类型
		fileType := http.DetectContentType(buffer)
		if fileType == "image/jpeg" || fileType == "image/jpg" ||
			fileType == "image/gif" || fileType == "image/png" {
			//保存图片
			err = c.SaveUploadedFile(file, fileName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				c.String(http.StatusOK, "upload successfully")
			}
		} else {
			//如果不是图片则不保存
			c.String(http.StatusOK, "file is not image")
		}


	})

	r.Run(":8080")
}