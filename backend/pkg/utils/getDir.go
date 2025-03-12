package utils // 获取程序当前工作路径
import (
	"log"
	"os"
)

func GetDir() string {
	dir, err := os.Getwd()
	// dir: "E:\\Film_Management_System\\backend"
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}

	// // debug
	// log.Printf("当前工作目录: %s", dir)
	return dir
}
