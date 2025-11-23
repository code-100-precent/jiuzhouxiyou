package main

import (
	"log"

	"lingecho-backend/config"
	"lingecho-backend/database"
	"lingecho-backend/routes"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 连接数据库
	database.Connect()

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	port := config.AppConfig.Port
	log.Printf("服务器启动在端口 %s", port)
	log.Printf("API 地址: http://localhost:%s/api/v1", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("服务器启动失败: ", err)
	}
}
