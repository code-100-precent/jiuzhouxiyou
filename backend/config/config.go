package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver  string
	DSN       string
	Port      string
	JWTSecret string
}

var AppConfig *Config

func LoadConfig() {
	// 尝试加载 .env 文件，如果不存在也不报错
	_ = godotenv.Load()

	AppConfig = &Config{
		DBDriver:  getEnv("DB_DRIVER", "mysql"),
		DSN:       getEnv("DSN", ""),
		Port:      getEnv("PORT", "7070"),
		JWTSecret: getEnv("JWT_SECRET", "default-secret-key"),
	}

	if AppConfig.DSN == "" {
		log.Fatal("DSN 配置不能为空")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
