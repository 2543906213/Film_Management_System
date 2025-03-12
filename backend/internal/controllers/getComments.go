package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 评论处理函数示例
func GetComments(c *gin.Context) {
	// 可通过查询参数过滤，例如 ?photo_id=1
	photoID := c.Query("photo_id")
	if photoID != "" {
		c.JSON(http.StatusOK, gin.H{"data": "照片 " + photoID + " 的评论列表"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "所有评论列表"})
	}
}
