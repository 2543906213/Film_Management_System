package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置
	config.LoadConfig()

	// 获取加载后的配置
	cfg := config.GetConfig()

	// 初始化数据库服务器连接（无数据库，只连接服务器）
	database.ConnectMySQL(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	// 获取数据库实例
	db := database.GetDB()

	// 执行 schema.sql 文件进行数据库初始化
	database.RunSQL(db, "schema.sql")

	// 获取底层的*sql.DB 对象，为了关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}
	defer sqlDB.Close() // 关闭数据库连接

	// 可选: 自动迁移数据库（根据模型生成表结构）
	// DB.AutoMigrate(&YourModel{}) // 例如: 自动迁移你的模型

	// 初始化路由,创建一个默认的Gin路由器
	r := gin.Default()

	// 允许访问 /uploads 目录下的文件
	r.Static("/uploads", "./uploads")

	// 不同端口可能会遇到 CORS (跨域资源共享) 问题，需要在 gin 里加上 CORS 处理
	// 允许所有来源（开发时可以这样，生产环境需要配置具体的前端域名）
	r.Use(cors.Default())

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	serverPort := fmt.Sprintf(":%d", cfg.Server.Port)
	err = r.Run(serverPort)

	log.Printf("服务启动成功：%s", serverPort)

	if err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
