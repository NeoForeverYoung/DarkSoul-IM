package config

import (
	"im-server/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// 使用 SQLite，数据将存储在 im.db 文件中
	DB, err = gorm.Open(sqlite.Open("im.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移数据库表
	DB.AutoMigrate(&models.User{})
} 