package controllers

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 照片卡处理函数示例
func GetPhotos(c *gin.Context) {
	// 此处应调用数据库获取所有照片卡数据（支持分页、排序等）
	var photos []models.PhotoCard
	// 调用全局变量 db ，连接到数据库
	db := database.GetDB()
	// 预加载 Tags 关联的数据
	if err := db.Preload("Tags").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, photos)
}
