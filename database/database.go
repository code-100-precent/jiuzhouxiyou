package database

import (
	"log"

	"lingecho-backend/config"
	"lingecho-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// fixTableStructure 修复表结构
func fixTableStructure(db *gorm.DB) {
	// 如果 users 表已存在，尝试修复 uid 列的唯一约束
	if db.Migrator().HasTable(&models.User{}) {
		// 检查 uid 列是否存在唯一索引
		var indexExists bool
		db.Raw(`
			SELECT COUNT(*) > 0 
			FROM information_schema.statistics 
			WHERE table_schema = DATABASE() 
			AND table_name = 'users' 
			AND column_name = 'uid' 
			AND non_unique = 0
		`).Scan(&indexExists)

		if !indexExists {
			log.Println("尝试修复 users 表的 uid 唯一约束...")
			// 先尝试删除可能存在的旧索引（MySQL 不支持 IF EXISTS，所以忽略错误）
			_ = db.Exec("ALTER TABLE users DROP INDEX uid")
			// 添加唯一索引
			if err := db.Exec("ALTER TABLE users ADD UNIQUE INDEX uid (uid)").Error; err != nil {
				log.Printf("添加唯一索引失败（可能已有约束）: %v", err)
			} else {
				log.Println("成功添加 uid 唯一索引")
			}
		}
	}
}

// cleanupInvalidData 清理无效数据
func cleanupInvalidData(db *gorm.DB) {
	// 检查 users 表是否存在
	if !db.Migrator().HasTable(&models.User{}) {
		return
	}

	// 检查 uid 列是否存在
	var uidColumnExists bool
	db.Raw(`
		SELECT COUNT(*) > 0 
		FROM information_schema.columns 
		WHERE table_schema = DATABASE() 
		AND table_name = 'users' 
		AND column_name = 'uid'
	`).Scan(&uidColumnExists)

	if !uidColumnExists {
		log.Println("users 表存在但 uid 列不存在，跳过数据清理")
		return
	}

	// 删除 uid 为空字符串的记录
	result := db.Exec("DELETE FROM users WHERE uid = '' OR uid IS NULL")
	if result.Error != nil {
		log.Printf("清理无效用户数据时出错: %v", result.Error)
	} else if result.RowsAffected > 0 {
		log.Printf("已清理 %d 条无效用户记录", result.RowsAffected)
	}

	// 检查是否有重复的 uid（除了空字符串）
	var duplicateUIDs []string
	db.Raw(`
		SELECT uid FROM users 
		WHERE uid != '' AND uid IS NOT NULL
		GROUP BY uid 
		HAVING COUNT(*) > 1
	`).Scan(&duplicateUIDs)

	if len(duplicateUIDs) > 0 {
		log.Printf("发现重复的 uid: %v，将保留第一条记录", duplicateUIDs)
		for _, uid := range duplicateUIDs {
			// 保留第一条，删除其他重复记录
			db.Exec(`
				DELETE FROM users 
				WHERE uid = ? AND id NOT IN (
					SELECT id FROM (
						SELECT MIN(id) as id FROM users WHERE uid = ?
					) as t
				)
			`, uid, uid)
		}
	}
}

var DB *gorm.DB

func Connect() {
	var err error

	dsn := config.AppConfig.DSN
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}

	log.Println("数据库连接成功")

	// 清理无效数据（uid 为空或重复的记录）
	cleanupInvalidData(DB)

	// 修复表结构（如果表已存在但结构不对）
	fixTableStructure(DB)

	// 自动迁移
	err = DB.AutoMigrate(&models.User{}, &models.Archive{})
	if err != nil {
		log.Fatal("数据库迁移失败: ", err)
	}

	log.Println("数据库迁移完成")
}
