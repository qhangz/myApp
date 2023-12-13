package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/myapp/model"
	"github.com/myapp/service"
)

// @Title			PublishDiscuss
// @Description		publish discuss
// @Param			author  	formData	string		true	"作者名"
// @Param			title   	formData	string		true	"标题"
// @Param			summary		formData	string		true	"摘要"
// @Param			content		formData	string		true	"内容"
// @Success			200			object		controllers.Response	"publish success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/poublish [post]
func PublishDiscuss(c *gin.Context) {
	// get discuss info from request
	newDiscuss := model.Discuss{
		Author:  c.PostForm("author"),
		Title:   c.PostForm("title"),
		Summary: c.PostForm("summary"),
		Content: c.PostForm("content"),
	}

	err := service.PublishDiscuss(newDiscuss)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "publish success",
	})
}

// @Title			GetDiscussInfo
// @Description		get discuss info by discuss id from request
// @Param			discussID  	formData	int		true	"discuss ID"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/GetDiscussInfo [get]
func GetDiscussInfo(c *gin.Context) {
	// 将discussID转换为无符号整数类型
	discussID, strconEerr := strconv.ParseUint(c.Query("discussID"), 10, 64)
	if strconEerr != nil {
		// 处理转换错误
		c.JSON(http.StatusOK, gin.H{
			"code":      "500",
			"error":     strconEerr.Error(),
			"discussid": c.Param("discussID"),
			"disID":     discussID,
		})
		return
	}
	thisDiscussID := uint(discussID)
	commentList, err := service.GetDiscussInfo(thisDiscussID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": commentList,
	})
}
