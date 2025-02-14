package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"im-server/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/im_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	
	// 自动迁移数据库表
	DB.AutoMigrate(&models.User{})
} 