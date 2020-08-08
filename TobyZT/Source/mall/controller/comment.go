package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
)

func GetComments(c *gin.Context) {
	var req struct{ ID string }
	err := c.BindJSON(&req)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	comments, err := model.GetComments(req.ID)
	var res []gin.H
	for i := 0; i < len(comments); i++ {
		res = append(res, gin.H{
			"username": comments[i].Username,
			"content":  comments[i].Content,
			"time":     comments[i].Time.String(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "data": res})

}

func AddComment(c *gin.Context) {
	var form model.CommentRequest
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	username, exist := c.Get("username")
	if !exist || form.Username != username {
		failMsg(c, http.StatusUnauthorized, "user not login")
		return
	}
	err = checkComment(form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	err = model.AddComment(form)
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "data": ""})
}

func checkComment(form model.CommentRequest) (err error) {

	return nil
}
