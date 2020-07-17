/* Router Package sets up route rules */
package router

import (
	"github.com/gin-gonic/gin"

	"account/controller"
)

// InitRouter creates a gin engine, sets up its route rules
// and then return the pointer of the engine
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("api/v1/login", controller.Login)
	router.POST("api/v1/signup", controller.Signup)
	router.GET("api/v1/users/all", controller.QueryAllUsers)
	router.GET("api/v1/users/:userid", controller.QueryOne)
	router.DELETE("api/v1/users/:userid", controller.Delete)
	router.PUT("api/v1/users/:userid", controller.Update)


	return router
}
