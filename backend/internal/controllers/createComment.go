package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	// 解析请求中的评论数据并保存到数据库
	c.JSON(http.StatusCreated, gin.H{"data": "新增评论"})
}
