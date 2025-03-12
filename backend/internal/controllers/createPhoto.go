package controllers

import (
	"backend/internal/database"
	"backend/internal/models"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 增: CreatePhoto 新增照片卡的处理函数
func CreatePhoto(c *gin.Context) {
	// 获取文本字段
	title := c.PostForm("title")
	description := c.PostForm("description")
	shootingDateStr := c.PostForm("shooting_date") // 格式应为 "YYYY-MM-DD"
	shootingLocation := c.PostForm("shooting_location")
	filmType := c.PostForm("film_type")
	camera := c.PostForm("camera")
	tagsStr := c.PostForm("tags") // 获取 tags 字段，假设以逗号分隔

	// 获取上传的文件，表单字段名为 "photoFile"
	file, err := c.FormFile("photoFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传照片文件"})
		return
	}

	// 解析 shooting_date 字符串为 time.Time 类型
	shootingDate, err := time.Parse("2006-01-02", shootingDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "日期格式错误，正确格式为 YYYY-MM-DD"})
		return
	}

	// 根据日期生成存储目录（例如：./uploads/2025/03/）
	year := shootingDate.Format("2006")
	month := shootingDate.Format("01")
	dirPath := filepath.Join("uploads", year, month)

	// 自动创建目录（如果不存在）
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建目录失败: " + err.Error()})
		return
	}

	// 定义保存文件的完整路径
	// dst := dirPath + "/" + file.Filename
	dst := filepath.Join(dirPath, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将路径中的 \ 替换为 /,并在前面添加/
	dst = "/" + strings.ReplaceAll(dst, `\`, `/`)

	// 创建 PhotoCard 实例，PhotoURL 字段保存文件存储路径（或相对 URL）
	photo := models.PhotoCard{
		Title:            title,
		Description:      description,
		PhotoURL:         dst,
		ShootingDate:     shootingDate,
		ShootingLocation: shootingLocation,
		FilmType:         filmType,
		Camera:           camera,
	}

	// 调用全局变量 db ，连接到数据库
	db := database.GetDB()

	// 处理 tags 字段
	if tagsStr != "" {
		// 将 tagsStr 按逗号分割为切片
		tags := strings.Split(tagsStr, ",")
		for _, tagName := range tags {
			tagName = strings.TrimSpace(tagName)
			if tagName == "" {
				continue
			}
			var tag models.Tag
			// 查找标签是否已存在
			if err := db.Where("name = ?", tagName).First(&tag).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// 如果标签不存在，则创建新标签
					tag = models.Tag{Name: tagName}
					if err := db.Create(&tag).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
			// 将标签添加到照片的 Tags 列表中
			photo.Tags = append(photo.Tags, tag)
		}
	}

	// 使用 GORM 将 photo 保存到数据库中
	if err := db.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回 201 Created 状态和新创建的照片数据
	c.JSON(http.StatusCreated, photo)

}
