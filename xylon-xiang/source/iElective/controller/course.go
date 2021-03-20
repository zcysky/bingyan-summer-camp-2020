package controller

import (
	"github.com/labstack/echo"
	"iElective/module"
	"iElective/util"
	"net/http"
	"strconv"
)

type SelectObject struct {
	Class         string
	Time          int64
	Place         string
	SigningMethod string
	examMethod    string
}

func SelectCourseClass(ctx echo.Context) error {

	var (
		selectObj SelectObject
		err       error
	)

	// get the query param, which are used to set the boundary
	selectObj.Class = ctx.QueryParam("class")
	time := ctx.QueryParam("time")
	selectObj.Place = ctx.QueryParam("place")
	selectObj.SigningMethod = ctx.QueryParam("signing_method")
	selectObj.examMethod = ctx.QueryParam("exam_method")

	if time != "" {
		selectObj.Time, err = strconv.ParseInt(time, 10, 64)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "time convert error")
		}
	}
	selectObj.Time = 0

	results, err := SelectService(selectObj)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "unknown error")
	}

	return ctx.JSON(http.StatusOK, results.([]module.Course))
}

func SelectSpecificCourse(ctx echo.Context) error {
	courseId := ctx.Param("courseId")

	result, err := SelectService(SelectObject{}, courseId)
	if err != nil {
		return ctx.String(http.StatusOK, "unknown error")
	}

	return ctx.JSON(http.StatusOK, result.(module.Course))
}

// return the []Course when courseId is nil
// else the interface is Course
func SelectService(obj SelectObject, courseId ...string) (interface{}, error) {

	if courseId != nil {
		result, err := module.SelectSpecificCourse(courseId[0])
		if err != nil {
			return module.Course{}, err
		}

		return result, nil
	}

	results, err := module.SelectCourseClass(obj)
	if err != nil {
		return module.Course{}, err
	}

	return results, nil
}

func PublicCourseInfo(ctx echo.Context) error {

	course := new(module.Course)

	if err := ctx.Bind(course); err != nil {
		return ctx.String(http.StatusInternalServerError, "Json bind failure")
	}

	courseId, err := PubCourseService(course)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "unknown error\n The courseId is "+courseId)
	}

	return ctx.String(http.StatusOK, "Insert into db successfully\n The courseId is "+courseId)
}

func PubCourseService(course *module.Course) (string, error) {

	uid, err := util.GenerateUUID()
	if err != nil {
		return "", err
	}

	course.CourseId = uid

	err = module.Insert(module.COURSE, *course)
	if err != nil {
		return uid, err
	}
	return uid, nil
}
