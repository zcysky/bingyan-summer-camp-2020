package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"iElective/controller"
)

func main() {

	e := echo.New()

	// user login
	e.GET("/log", controller.Log)

	// public the course
	e.POST("/course", controller.PublicCourseInfo, middleware.JWTWithConfig(controller.JwtConfig))

	// search the kind of the courses
	e.GET("/course/", controller.SelectCourseClass, middleware.JWTWithConfig(controller.JwtConfig))

	// search the specific course
	e.GET("/course/:courseId", controller.SelectSpecificCourse, middleware.JWTWithConfig(controller.JwtConfig))

	// public the comment to the course or the user
	e.POST("/comment/:courseId", controller.PubComment, middleware.JWTWithConfig(controller.JwtConfig))

	// the user search his comment or who comment his comment
	e.GET("/comment", controller.GetComment, middleware.JWTWithConfig(controller.JwtConfig))

	e.Logger.Fatal(e.Start("1323"))
}
