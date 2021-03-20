package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"iElective/module"
	"iElective/util"
	"net/http"
)

func PubComment(ctx echo.Context) error {

	comment := new(module.UserComment)

	if err := ctx.Bind(comment); err != nil {
		return util.HttpErrorHandle(ctx, err)
	}

	commentId, err := PubCommentService(comment)

	if err != nil {
		return util.HttpErrorHandle(ctx, err)
	}

	return ctx.String(http.StatusOK, "PubComment successfully\n The comment id is "+commentId)

}

func GetComment(ctx echo.Context) error {

	userId := ctx.QueryParam("userId")
	targetId := ctx.QueryParam("targetId")

	JsonResults, err := GetCommentService(userId, targetId)
	if err != nil {
		return util.HttpErrorHandle(ctx, err)
	}

	return ctx.JSON(http.StatusOK, JsonResults)
}

func PubCommentService(comment *module.UserComment) (string, error) {

	commentId, err := util.GenerateUUID()
	if err != nil {
		return "", err
	}

	comment.CommentId = commentId

	// just a comment for a course
	if comment.TargetId == "" {

		courseComment := comment.CourseComment

		err = module.Insert(module.COURSECOMMENT, courseComment)
		if err != nil {
			return commentId, err
		}

		return commentId, nil
	}

	err = module.Insert(module.USERCOMMENT, *comment)
	if err != nil {
		return commentId, err
	}

	return commentId, nil

}

func GetCommentService(userId string, targetId string) ([]byte, error) {

	// just query the userId for CourseComment
	if targetId == "" {
		results, err := module.SelectComment(module.COURSECOMMENT, userId)
		if err != nil {
			return nil, err
		}

		results = results.([]module.CourseComment)

		byteStream, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return byteStream, nil
	}

	// else query the targetId for userComment

	results, err := module.SelectComment(module.USERCOMMENT, targetId)
	if err != nil {
		return nil, err
	}

	results = results.([]module.UserComment)

	byteStream, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return byteStream, nil

}
