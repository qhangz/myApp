package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/myApp/model"
	"github.com/myApp/service"
)

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": user,
	})
}

// @Title			user register
// @Description		使用表单用户名、密码和邮箱注册
// @Param			username	formData	string		true	"用户名"
// @Param			password	formData	string		true	"密码"
// @Param			email		formData	string		true	"邮箱账号"
// @Success			200			object		controllers.Response	"register success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/user/register [post]
func Register(c *gin.Context) {
	// get user info from request
	newUser := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email:    c.PostForm("email"),
	}

	// c.Bind(&newUser)
	err := service.Register(newUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "register success",
	})
}

// @Title			user login
// @Description		使用表单用户名和密码登录
// @Param			username	formData	string		true	"用户名"
// @Param			password	formData	string		true	"密码"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/user/login [post]
func Login(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	// c.Bind(&user)
	userInfo, token, err := service.Login(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  "500",
			"error": "token is empty",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"token": token,
			"userInfo": gin.H{
				"username":     userInfo.Username,
				"age":          userInfo.Age,
				"summary":      userInfo.Summary,
				"avatar_image": userInfo.Avatar_image,
				"created_at":   userInfo.CreatedAt,
			},
		},
		"msg": "登录成功",
	})
}

// @Title			user list
// @Description		获取用户列表信息
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/user/list [get]
func GetUserList(c *gin.Context) {
	userList, err := service.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": userList,
	})
}
