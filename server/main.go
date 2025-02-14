package main

import (
	"im-server/config"
	"im-server/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	// 初始化数据库
	config.InitDB()
	
	// 设置路由
	r := routes.SetupRouter()
	
	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 启动服务器
	r.Run(":8080")
} 