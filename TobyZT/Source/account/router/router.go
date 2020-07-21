/* Router Package sets up route rules  */
package router

import (
	"account/controller"
	"github.com/gin-gonic/gin"
)

// InitRouter creates a gin engine, sets up its route rules
// and then return the pointer of the engine
func InitRouter() *gin.Engine {
	router := gin.Default()
	userGroup := router.Group("/api/v1")
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/signup", controller.Signup)
		userGroup.PUT("/users", controller.Update)
	}

	adminGroup := router.Group("/api/v1")
	{
		adminGroup.GET("/users", controller.QueryAllUsers)
		adminGroup.GET("/users/:userid", controller.QueryOne)
		adminGroup.DELETE("/users/:userid", controller.Delete)
	}

	return router
}
