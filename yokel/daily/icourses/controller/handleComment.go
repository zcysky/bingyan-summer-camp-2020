package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"icourses/model"
	"icourses/defination"
	"icourses/config"
	"net/http"
	"strconv"
)

func convertFormToComment(form defination.CommentForm,lid int)defination.Comment{
	var comment defination.Comment
	comment.Uid=form.Uid
	comment.Time=form.Time
	comment.Lid=lid
	comment.Content=form.Content
	comment.Icnt=form.Icnt
	comment.Img=form.Img
	comment.TermEnd=form.TermEnd
	comment.User=form.User
	comment.Avatar=form.Avatar
	return comment
}

func HandlePostComment(ctx echo.Context)error{
	//直接发表评论较为特殊，实际包含了评论和评价两方面
	//atdc,exam,evaluation属于评价
	var form defination.CommentForm
	err:=ctx.Bind(&form)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"无法解析参数")
	}
	lidStr:=ctx.Param("lid")
	lid,err:=strconv.Atoi(lidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	comment:=convertFormToComment(form,lid)
	config.Config.MongoConfig.Cid++
	comment.Cid=config.Config.MongoConfig.Cid
	//userInfo,err:=model.FindUserWithUid(form.Uid)
	//if err!=nil {
	//	if err==mongo.ErrNoDocuments{
	//		return ctx.String(http.StatusBadRequest,"数据库查询无该用户")
	//	}
	//	return ctx.String(http.StatusInternalServerError,"数据库查询错误")
	//}
	//comment.User=userInfo.User
	//comment.Avatar=userInfo.Avatar
	err=model.InsertComment(comment)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"数据库查询错误")
	}
	return ctx.String(http.StatusOK,"发表评论成功")
}

func HandlePostSubcmt(ctx echo.Context)error{
	var form defination.SubComment
	err:=ctx.Bind(&form)
	cidStr:=ctx.Param("cid")
	cid,err:=strconv.Atoi(cidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"无法解析参数")
	}
	config.Config.MongoConfig.Cid++
	form.Cid=config.Config.MongoConfig.Cid
	//comment,err:=model.FindCommentWithCid(cid)
	err=model.InsertSubCmt(cid,form)
	if err !=nil{
		return ctx.String(http.StatusInternalServerError,"数据库更新错误")
	}
	return ctx.String(http.StatusOK,"succeed")
}

func HandleGetAllComments(ctx echo.Context)error{
	lidStr:=ctx.Param("lid")
	lid,err:=strconv.Atoi(lidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	commentInfo,err:=model.FindCommentAll(lid)
	if err!=nil {
		return ctx.String(http.StatusBadRequest,"无法解析参数")
	}
	commentInfoJson,err:=json.Marshal(commentInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"将数据转为JSON时错误")
	}
	return ctx.JSON(http.StatusOK,commentInfoJson)
}

func HandleGetComment(ctx echo.Context)error{
	cidStr:=ctx.Param("cid")
	cid,err:=strconv.Atoi(cidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	commentInfo,err:=model.FindCommentWithCid(cid)
	if err!=nil {
		return ctx.String(http.StatusBadRequest,"无法解析参数")
	}
	commentInfoJson,err:=json.Marshal(commentInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"将数据转为JSON时错误")
	}
	return ctx.JSON(http.StatusOK,commentInfoJson)

}