package controllers

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 改
func UpdatePhoto(c *gin.Context) {
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

	// 获取表单字段而不是JSON
	title := c.PostForm("title")
	description := c.PostForm("description")
	shootingDateStr := c.PostForm("shooting_date")
	shootingLocation := c.PostForm("shooting_location")
	filmType := c.PostForm("film_type")
	camera := c.PostForm("camera")
	tagsStr := c.PostForm("tags") // 标签以逗号分隔

	// 更新标准字段
	if title != "" {
		photoCard.Title = title
	}
	if description != "" {
		photoCard.Description = description
	}
	if shootingLocation != "" {
		photoCard.ShootingLocation = shootingLocation
	}
	if filmType != "" {
		photoCard.FilmType = filmType
	}
	if camera != "" {
		photoCard.Camera = camera
	}

	// 处理日期
	if shootingDateStr != "" {
		shootingDate, err := time.Parse("2006-01-02", shootingDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "日期格式错误，正确格式为 YYYY-MM-DD"})
			return
		}
		photoCard.ShootingDate = shootingDate
	}

	// 处理文件上传（如果有新文件）
	file, err := c.FormFile("photoFile")
	if err == nil { // 没有错误意味着有文件上传
		// 根据日期生成存储目录
		year := photoCard.ShootingDate.Format("2006")
		month := photoCard.ShootingDate.Format("01")
		dirPath := filepath.Join("uploads", year, month)

		// 自动创建目录（如果不存在）
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建目录失败: " + err.Error()})
			return
		}

		// 如果之前有图片，删除旧文件
		if photoCard.PhotoURL != "" {
			oldFilePath := strings.TrimPrefix(photoCard.PhotoURL, "/")
			if _, err := os.Stat(oldFilePath); err == nil {
				os.Remove(oldFilePath) // 尝试删除旧文件，忽略可能的错误
			}
		}

		// 保存新文件
		dst := filepath.Join(dirPath, file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败: " + err.Error()})
			return
		}

		// 更新照片URL
		photoCard.PhotoURL = "/" + strings.ReplaceAll(dst, `\`, `/`)
	}

	// 开启事务处理标签关系更新
	tx := database.DB.Begin()

	// 如果有提供标签数据，则更新标签关系
	if tagsStr != "" {
		// 清除照片与所有标签之间的关联关系
		association := tx.Model(&photoCard).Association("Tags")
		err := association.Clear()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "清除标签关系失败"})
			return
		}

		// 处理新的标签关系
		var newTags []models.Tag
		tagNames := strings.Split(tagsStr, ",")
		for _, tagName := range tagNames {
			tagName = strings.TrimSpace(tagName)
			if tagName == "" {
				continue
			}

			// 检查标签是否存在
			var existingTag models.Tag
			if err := tx.Where("name = ?", tagName).First(&existingTag).Error; err == nil {
				// 标签已存在，使用现有标签
				newTags = append(newTags, existingTag)
			} else {
				// 标签不存在，创建新标签
				newTag := models.Tag{Name: tagName}
				if err := tx.Create(&newTag).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "创建新标签失败"})
					return
				}
				newTags = append(newTags, newTag)
			}
		}

		// 添加新的标签关系
		if err := tx.Model(&photoCard).Association("Tags").Replace(newTags); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新标签关系失败"})
			return
		}
	}

	// 将更新后的对象保存回数据库
	if err := tx.Save(&photoCard).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新照片卡失败: " + err.Error()})
		return
	}

	// 提交事务
	tx.Commit()

	// 返回更新后的照片卡数据
	c.JSON(http.StatusOK, gin.H{"message": "更新照片卡成功", "data": photoCard})
}
