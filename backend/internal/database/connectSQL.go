package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个全局数据库变量
var DB *gorm.DB

// 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// // 连接MySQL服务器，但目前还没有数据库
// func ConnectSQL(user, password, host string, port int) *sql.DB {
// 	// 构造服务器连接字符串（不指定数据库）
// 	cfg := mysql.NewConfig()
// 	cfg.User = user
// 	cfg.Passwd = password
// 	cfg.Addr = fmt.Sprintf("%s:%d", host, port)

// 	DSN := cfg.FormatDSN()

// 	// 打开数据库连接
// 	DB, err := sql.Open("mysql", DSN)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 测试连接是否成功
// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("连接到 MySQL 服务器失败: ", err)
// 		DB.Close()
// 		return nil
// 	}

// 	log.Println("数据库服务器连接成功!")
// 	return DB
// }

// 连接 MySQL 服务器（不指定具体数据库）
func ConnectMySQL(user, password, host string, port int) *gorm.DB {
	// 构造 DSN，不指定数据库名，附带常用参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port)

	// 使用 Gorm 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接到 MySQL 服务器失败: ", err)
	}

	// 获取底层 *sql.DB 对象，并检测连接是否成功
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("MySQL 服务器 Ping 失败: ", err)
	}

	log.Println("MySQL 服务器连接成功!")
	DB = db
	return db
}
