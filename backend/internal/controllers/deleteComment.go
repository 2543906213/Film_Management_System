package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	// 删除指定 id 的评论
	c.JSON(http.StatusOK, gin.H{"data": "删除评论 " + id})
}
