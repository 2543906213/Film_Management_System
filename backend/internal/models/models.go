package models

import (
	"time"
)

// PhotoCard 表示一张照片卡
type PhotoCard struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Title            string    `gorm:"type:varchar(255);not null" json:"title"`
	Description      string    `json:"description"`
	PhotoURL         string    `gorm:"not null" json:"photo_url"`
	ShootingDate     time.Time `gorm:"type:date" json:"shooting_date"`             // 拍摄日期
	ShootingLocation string    `gorm:"type:varchar(255)" json:"shooting_location"` // 拍摄地点
	FilmType         string    `gorm:"type:varchar(255)" json:"film_type"`         // 胶片类型
	Camera           string    `gorm:"type:varchar(255)" json:"camera"`            // 相机
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	Tags             []Tag     `gorm:"many2many:photo_card_tags;foreignKey:ID;joinForeignKey:PhotoCardID;References:ID;joinReferences:TagID" json:"tags"` // 多对多关联标签
}

// Tag 表示一个标签
type Tag struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Name       string      `gorm:"type:varchar(100);unique;not null" json:"name"`
	PhotoCards []PhotoCard `gorm:"many2many:photo_card_tags;foreignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:PhotoCardID" json:"photo_cards"` // 反向多对多关联
}

// 如果需要自定义中间表结构，可以定义 PhotoCardTag，但如果中间表不包含额外字段，GORM 可自动生成
type PhotoCardTag struct {
	PhotoID uint `gorm:"primaryKey"`
	TagID   uint `gorm:"primaryKey"`
}

// Comment 留言评论模型
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PhotoID   uint      `json:"photo_id"` // 如果评论关联某张照片
	UserName  string    `json:"user_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
