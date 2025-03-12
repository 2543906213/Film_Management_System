package controllers

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 删
func DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为整数
	photoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	// 调用全局变量 db ，连接到数据库
	db := database.GetDB()

	var photoCard models.PhotoCard
	// 预加载标签，以便后续可能的处理
	if err := db.Preload("Tags").First(&photoCard, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "照片未找到"})
		return
	}
	// 1. 清除与标签的关联关系
	if err := db.Model(&photoCard).Association("Tags").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "清除标签关系失败"})
		return
	}

	// 2. 删除照片记录
	if err := db.Delete(&photoCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除照片记录失败"})
		return
	}

	// 3. 删除物理文件
	if photoCard.PhotoURL != "" {
		// 1. 去掉URL前缀斜杠
		photoPath := strings.TrimPrefix(photoCard.PhotoURL, "/")

		// 2. 将路径转换为当前操作系统的格式（特别是Windows）
		photoPath = filepath.FromSlash(photoPath)

		if _, err := os.Stat(photoPath); err == nil {
			// 文件存在，尝试删除
			if err := os.Remove(photoPath); err != nil {
				// 只记录错误，不影响API响应结果
				// log.Printf("删除文件失败: %v, 路径: %s", err, photoPath)
			}
		}
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "照片删除成功",
		"id":      photoID,
	})
}
