package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"icourses/defination"
	"icourses/model"
	"net/http"
	"strconv"
)

func HandleGetAllCourses(ctx echo.Context)error{
	var filter defination.LessonFilter
	err:=ctx.Bind(&filter)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"无法解析参数")
	}
	lessonInfo,err:=model.FindAllCourseWithFilter(filter)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"数据库查询错误")
	}
	lessonInfoJson,err:=json.Marshal(lessonInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"课程数据转为JSON时错误")
	}
	return ctx.JSON(http.StatusOK,lessonInfoJson)
}

func HandleGetCourse(ctx echo.Context)error{
	lidStr:=ctx.QueryParam("lid")
	lid,err:=strconv.Atoi(lidStr)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"参数格式错误")
	}
	lessonInfo,err:=model.FindCourseWithLid(lid)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"数据库查询错误")
	}
	lessonInfoJson,err:=json.Marshal(lessonInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"课程数据转为JSON时错误")
	}
	return ctx.JSON(http.StatusOK,lessonInfoJson)
}