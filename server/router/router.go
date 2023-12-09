package router

import (
	// "os/user"

	"github.com/gin-gonic/gin"
	"github.com/myApp/controller"
	"github.com/myApp/middlewares"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(gin.DebugMode)
	// router := gin.New()
	router := gin.Default()
	router.Use(middlewares.Cors())

	// router.GET("/users/:username", controller.GetUserByUsername)

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/list", controller.GetUserList)
		user.GET("/info/:username", controller.GetUserByUsername)
		
	}

	return router
}
