package controllers

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查
// GetPhotoByID 根据ID获取照片卡详细信息
func GetPhotoByID(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为整数
	photoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	var photoCard models.PhotoCard
	// 预加载标签
	// 1. 准备查询 - 指定要预加载标签数据
	query := database.DB.Preload("Tags")

	// 2. 执行查询 - 查找特定ID的照片
	err = query.First(&photoCard, photoID).Error

	// 3. 处理可能的错误
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "照片未找到"})
		return
	}

	// 直接返回完整对象，包含了标签数据
	c.JSON(http.StatusOK, photoCard)
}
