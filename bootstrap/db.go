package bootstrap

import (
	"lichmaker/girlfriend-gift-1/app/models/schedule"

	"lichmaker/girlfriend-gift-1/pkg/config"
	"lichmaker/girlfriend-gift-1/pkg/model"
	"time"

	"gorm.io/gorm"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.Viper.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.Viper.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Viper.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&schedule.Schedule{},
	)
}
