package controller

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	file, err := c.FormFile("file")
	filename := username.(string) + "_" + c.PostForm("name")
	if err != nil {
		failMsg(c, http.StatusBadRequest, "upload failed")
		return
	}
	// check filename
	ptn := `^[a-zA-Z0-9_-]{1,12}(.jpg|.png|.bmp)$`
	reg := regexp.MustCompile(ptn)
	valid := reg.MatchString(filename)
	if !valid {
		failMsg(c, http.StatusBadRequest, "invalid file name")
		return
	}

	// save to local directory

	dst := "tmp/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		failMsg(c, http.StatusBadRequest, "upload failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"url": "pics/" + filename,
	})
}

func GetPicture(c *gin.Context) {
	filename := c.Param("filename")
	c.File("tmp/" + filename)
}
