package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/hw", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/rand/:seed", func(c *gin.Context) {
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
		fileName := c.PostForm("name")
		file, err := c.FormFile("image")
		fmt.Println(fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		buffer := make([]byte, 512)
		_, err = f.Read(buffer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		fileType := http.DetectContentType(buffer)
		if fileType == "image/jpeg" || fileType == "image/jpg" ||
			fileType == "image/gif" || fileType == "image/png" {
			err = c.SaveUploadedFile(file, fileName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				c.String(http.StatusOK, "upload successfully")
			}
		} else {
			c.String(http.StatusOK, "file is not image")
		}


	})

	r.Run(":8080")
}