package main
import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"icourses/config"
	"icourses/controller"
)
func main(){
	e:=echo.New()
	fmt.Println(config.Config.EchoConfig)
	e.GET(config.Config.EchoConfig.GetToken,controller.HandleGetAllCourses)
	e.GET(config.Config.EchoConfig.GetAllComments,controller.HandleGetAllComments)
	e.GET(config.Config.EchoConfig.GetComment,controller.HandleGetComment)
	e.GET(config.Config.EchoConfig.GetAllCourses,controller.HandleGetAllCourses)
	e.GET(config.Config.EchoConfig.GetCourse,controller.HandleGetCourse)

	afterLogin := e.Group("/user", middleware.JWT([]byte(config.Config.JwtConfig.Secret)))
	afterLogin.POST(config.Config.EchoConfig.PostComment,controller.HandlePostComment)
	afterLogin.POST(config.Config.EchoConfig.PostSubcmt,controller.HandlePostSubcmt)
	afterLogin.POST(config.Config.EchoConfig.PostUser,controller.HandlePostUser)
	afterLogin.POST(config.Config.EchoConfig.GetUserAllComments,controller.HandleGetUserAllComments)
	e.Start(config.Config.EchoConfig.EchoPort)
}