package main

import (
	"fmt"
	"log"

	"lingecho-backend/config"
	"lingecho-backend/database"
	"lingecho-backend/models"
	"lingecho-backend/utils"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 连接数据库
	database.Connect()

	// 初始化用户列表
	users := []struct {
		UID    string
		Pwd    string
		Name   string
		Avatar string
	}{
		{"test", "123456", "测试用户", "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png"},
		{"wukong", "123456", "孙悟空", "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png"},
		{"bajie", "123456", "猪八戒", "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar02.png"},
		{"wujing", "123456", "沙悟净", "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png"},
		{"tangseng", "123456", "唐僧", "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar02.png"},
	}

	for _, u := range users {
		// 检查用户是否已存在
		var existingUser models.User
		if err := database.DB.Where("uid = ?", u.UID).First(&existingUser).Error; err == nil {
			fmt.Printf("用户 %s 已存在，跳过\n", u.UID)
			continue
		}

		// 加密密码
		hashedPassword, err := utils.HashPassword(u.Pwd)
		if err != nil {
			log.Printf("加密密码失败: %v", err)
			continue
		}

		// 创建用户
		user := models.User{
			UID:      u.UID,
			Password: hashedPassword,
			Name:     u.Name,
			Avatar:   u.Avatar,
		}

		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("创建用户 %s 失败: %v", u.UID, err)
			continue
		}

		fmt.Printf("✓ 成功创建用户: %s (%s)\n", u.UID, u.Name)
	}

	fmt.Println("\n初始化完成！")
	fmt.Println("测试账号列表：")
	for _, u := range users {
		fmt.Printf("  账号: %s, 密码: %s, 昵称: %s\n", u.UID, u.Pwd, u.Name)
	}
}
