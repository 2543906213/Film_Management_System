package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {

	LoadConfig()

	// 获取配置并验证
	config := GetConfig()
	if config.Server.Port != 8080 {
		t.Errorf("期望 server.port 为 8080, 但得到了 %d", config.Server.Port)
	}
	if config.Database.Host != "localhost" {
		t.Errorf("期望 database.host 为 'localhost', 但得到了 %s", config.Database.Host)
	}
}
