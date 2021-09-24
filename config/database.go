package config

import (
	"lichmaker/girlfriend-gift-1/pkg/config"
)

func init() {
	config.Viper.Set("database", map[string]interface{}{
		"mysql": map[string]interface{}{
			// 数据库连接信息
			"host":     config.Get("DB_HOST", "127.0.0.1"),
			"port":     config.Get("DB_PORT", "3306"),
			"database": config.Get("DB_DATABASE", ""),
			"username": config.Get("DB_USERNAME", ""),
			"password": config.Get("DB_PASSWORD", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Get("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Get("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Get("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}
