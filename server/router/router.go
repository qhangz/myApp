package router

import (
	// "os/user"

	"github.com/gin-gonic/gin"
	"github.com/myapp/controller"
	"github.com/myapp/middlewares"
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
		user.POST("/update/username", controller.UpdateUsername)
		user.POST("/update/password", controller.UpdatePassword)
		user.POST("/update/email", controller.UpdateEmail)
		user.POST("/update/age", controller.UpdateAge)
		user.POST("/update/summary", controller.UpdateSummary)
		user.POST("/delete", controller.DeleteUser)
	}

	discuss := router.Group("/api/discuss")
	{
		discuss.POST("/publish", controller.PublishDiscuss)
		discuss.GET("/discussinfo", controller.GetDiscussInfo)
	}

	comment := router.Group("/api/comment")
	{
		comment.POST("/publish", controller.PublishComment)
	}

	return router
}
