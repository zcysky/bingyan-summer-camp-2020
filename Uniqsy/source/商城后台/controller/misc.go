package controller

//出错时返回信息
import (
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

//失败时返回信息
func fail(c *gin.Context, status int, err string)  {
	c.JSON(status, gin.H{
		"success":	false,
		"error":	err,
		"data":		"",
	})
	c.Abort()
}

//成功时返回信息
func successStr(c *gin.Context, status int, data string) {
	c.JSON(status, gin.H{
		"success":	true,
		"error":	"",
		"data":		data,
	})
}

func successH(c *gin.Context, status int, data gin.H) {
	c.JSON(status, gin.H{
		"success":	true,
		"error":	"",
		"data":		data,
	})
}

func successHList(c *gin.Context, status int, data []gin.H) {
	c.JSON(status, gin.H{
		"success":	true,
		"error":	"",
		"data":		data,
	})
}

func successStrList(c *gin.Context, status int, data []string) {
	c.JSON(status, gin.H{
		"success":	true,
		"error":	"",
		"data":		data,
	})
}

//保存上传的图片
func UploadPics(c *gin.Context) {
	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//获取文件
	file, err := c.FormFile("file")
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//校验文件格式
	err = checkFile(file)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//获取文件名
	fileName := c.PostForm("name")

	//确定文件保存位置
	dest := "userFile/" + userName.(string) + "/pics/" + fileName

	//保存文件
	err = c.SaveUploadedFile(file, dest)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	data := gin.H{
		"url":	dest,
	}
	successH(c, http.StatusOK, data)
}

func checkFile(file *multipart.FileHeader) (err error) {
	//打开图片，
	f, err := file.Open()
	if err != nil {
		return err
	}

	//读取图片的前512个字节，用于鉴定文件是否为图片
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	//鉴定文件类型
	fileType := http.DetectContentType(buffer)
	if fileType == "image/jpeg" || fileType == "image/jpg" ||
		fileType == "image/gif" || fileType == "image/png" {
		return nil
	} else {
		return errors.New("file is not a pics")
	}
}