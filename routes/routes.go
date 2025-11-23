package routes

import (
	"lingecho-backend/handlers"
	"lingecho-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API v1
	v1 := r.Group("/api/v1")
	{
		// 用户相关（不需要认证）
		users := v1.Group("/users")
		{
			users.POST("/login", handlers.Login)
			users.POST("/register", handlers.Register)
		}

		// 需要认证的路由
		auth := v1.Group("/users")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/ping", handlers.Ping)
			auth.GET("/archive", handlers.GetArchives)
			auth.POST("/archive", handlers.CreateArchive)
		}
	}

	return r
}
