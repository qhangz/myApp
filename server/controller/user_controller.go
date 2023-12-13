package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/myapp/model"
	"github.com/myapp/service"
)

// @Title			get user info by username from request
// @Description		通过用户名获取用户信息
// @Param			username	formData	string		true	"用户名"
// @Success			200			object		controllers.Response	"register success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/info/:username [get]
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
// @Router			/api/user/register [post]
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
// @Router			/api/user/login [post]
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
// @Router			/api/user/list [get]
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

// @Title			user info update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update [post]
// func UpdateUserInfo(c *gin.Context) {
// 	// get user info from request
// 	user := model.User{
// 		Username:     c.PostForm("username"),
// 		Password:     c.PostForm("password"),
// 		Email:        c.PostForm("email"),
// 		Age:          0,
// 		Summary:      c.PostForm("summary"),
// 		Avatar_image: c.PostForm("avatar_image"),
// 	}
// 	// 将字符串类型的年龄转换为无符号整数类型
// 	age, strconvErr := strconv.ParseUint(c.PostForm("age"), 10, 64)
// 	if strconvErr != nil {
// 		// 处理转换错误
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":  "500",
// 			"error": strconvErr.Error(),
// 		})
// 		return
// 	}
// 	user.Age = int(age)
// err := service.UpdateUserInfo(user)
// if err != nil {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":  "400",
// 		"error": err.Error(),
// 	})
// 	return
// }
// c.JSON(http.StatusOK, gin.H{
// 	"code": "200",
// 	"msg":  "update success",
// })
// }

// @Title			user username update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Param			newusername	formData	string		true	"新用户名"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update/username [post]
func UpdateUsername(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
	}
	newUsername := c.PostForm("newusername")

	err := service.UpdateUsername(user, newUsername)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "update username success",
	})
}

// @Title			user password update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Param			password	formData	string		true	"密码"
// @Param			newpassword	formData	string		true	"新密码"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update/password [post]
func UpdatePassword(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	newPassword := c.PostForm("newpassword")

	err := service.UpdatePassword(user, newPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "update password success",
	})
}

// @Title			user email update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Param			newemail	formData	string		true	"新邮箱"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update/emali [post]
func UpdateEmail(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
	}
	newEmail := c.PostForm("newemail")

	err := service.UpdateEmail(user, newEmail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "update email success",
	})
}

// @Title			user age update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Param			newage   	formData	string		true	"新年龄"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update/age [post]
func UpdateAge(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
	}
	tempAge, strconvErr := strconv.ParseUint(c.PostForm("newage"), 10, 64)
	if strconvErr != nil {
		// 处理转换错误
		c.JSON(http.StatusOK, gin.H{
			"code":  "500",
			"error": strconvErr.Error(),
		})
		return
	}
	newAge := int(tempAge)

	err := service.UpdateAge(user, newAge)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "update age success",
	})
}

// @Title			user summary update
// @Description		使用用户名修改用户信息(username,password,email,age,summary,avatar_image)
// @Param			username	formData	string		true	"用户名"
// @Param			newsummary  formData	string		true	"新简介"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/update/summary [post]
func UpdateSummary(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
	}
	newSummary := c.PostForm("newsummary")

	err := service.UpdateSummary(user, newSummary)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "update summary success",
	})
}

// @Title			user delete
// @Description		使用用户名删除用户信息
// @Param			username	formData	string		true	"用户名"
// @Success			200			object		controllers.Response	"code,data,msg"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/user/delete [post]
func DeleteUser(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
	}

	err := service.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "delete user success",
	})
}