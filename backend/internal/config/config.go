package config

import (
	"backend/pkg/utils"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Database struct {
		User         string `json:"user"`
		Password     string `json:"password"`
		Host         string `json:"host"`
		Port         int    `json:"port"`
		DatabaseName string `json:"name"`
	} `json:"database"`
}

var AppConfig Config

func GetConfig() *Config {
	return &AppConfig
}

// 加载配置文件
func LoadConfig() {
	// 获取程序当前工作路径
	dir := utils.GetDir()

	// 得到同文件夹下的json文件路径
	configPath := filepath.Join(dir, "internal", "config", "config.json")

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析 JSON 数据
	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	// 打印配置信息
	log.Printf("加载配置成功: %+v", AppConfig)
}
