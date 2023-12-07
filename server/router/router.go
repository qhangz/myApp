package router

import (
	// "os/user"

	"github.com/myApp/controller"
	"github.com/myApp/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(gin.DebugMode)
	// router := gin.New()
	router := gin.Default()
	router.Use(middlewares.Cors())

	router.GET("/users/:username", controller.GetUserByUsername)

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	return router
}
