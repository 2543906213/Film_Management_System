package database

import (
	"backend/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

// RunSQL 读取并执行 sql 文件中的所有 SQL 语句
func RunSQL(db *gorm.DB, sqlName string) {
	// 获取程序当前工作路径
	dir := utils.GetDir()

	// 得到同文件夹下的sql文件路径
	sqlPath := filepath.Join(dir, "internal", "database", sqlName)

	data, err := os.ReadFile(sqlPath)
	if err != nil {
		log.Fatalf("读取 %s 文件失败: %v", sqlPath, err)
	}

	// 将文件内容按分号分割为多个语句
	statements := strings.Split(string(data), ";")

	// 执行每个语句
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		result := db.Exec(stmt)
		if result.Error != nil {
			log.Fatalf("执行 SQL 语句失败: %s\n错误: %v", stmt, err)
		}
	}

	log.Printf("执行 %s 文件成功! \n ", sqlName)
}
