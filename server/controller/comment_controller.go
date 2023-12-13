package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/myapp/model"
	"github.com/myapp/service"
)

// @Title			PublishComment
// @Description		publish comment
// @Param			author  	formData	string		true	"作者名"
// @Param			discussID	formData	int 		true	"discussID"
// @Param			content		formData	string		true	"内容"
// @Success			200			object		controllers.Response	"publish success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/poublish [post]
func PublishComment(c *gin.Context) {
	// get discuss info from request
	// newComment := model.Comment{
	// 	Author:    c.PostForm("author"),
	// 	// DiscussID: c.PostForm("discussID"),
	// 	Content:   c.PostForm("content"),
	// }
	newComment := model.Comment{
		Author:    c.PostForm("author"),
		DiscussID: 0,
		Content:   c.PostForm("content"),
		// DiscussID: c.PostForm("discussID")
	}
	// 将字符串类型的评论ID转换为无符号整数类型
	discussID, err := strconv.ParseUint(c.PostForm("discussid"), 10, 64)
	if err != nil {
		// 处理转换错误
		c.JSON(http.StatusOK, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	newComment.DiscussID = uint(discussID)

	publishErr := service.PublishComment(newComment)
	if publishErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": publishErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "publish success",
	})
}
