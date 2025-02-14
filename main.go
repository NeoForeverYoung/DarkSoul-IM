package main

import (
	"im-server/config"
	"im-server/routes"
)

func main() {
	// 初始化数据库
	config.InitDB()
	
	// 设置路由
	r := routes.SetupRouter()
	
	// 启动服务器
	r.Run(":8080")
} 