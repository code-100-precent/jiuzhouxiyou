package handlers

import (
	"net/http"
	"time"

	"lingecho-backend/database"
	"lingecho-backend/models"

	"github.com/gin-gonic/gin"
)

type ArchiveResponse struct {
	Archives []models.Archive `json:"archives"`
}

// GetArchives 获取用户存档列表
func GetArchives(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "未授权",
		})
		return
	}

	var archives []models.Archive
	if err := database.DB.Where("user_id = ?", userID).
		Order("last_login_time DESC").
		Limit(8).
		Find(&archives).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取存档列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": ArchiveResponse{
			Archives: archives,
		},
	})
}

type CreateArchiveRequest struct {
	Name string `json:"name" binding:"required"`
	Job  string `json:"job" binding:"required"`
}

// CreateArchive 创建新存档
func CreateArchive(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "未授权",
		})
		return
	}

	var req CreateArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误: " + err.Error(),
		})
		return
	}

	// 检查用户是否已有8个存档
	var count int64
	if err := database.DB.Model(&models.Archive{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询存档数量失败",
		})
		return
	}

	if count >= 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "存档数量已达上限（8个）",
		})
		return
	}

	// 创建新存档
	archive := models.Archive{
		UserID:        userID.(uint),
		Name:          req.Name,
		Job:           req.Job,
		Avatar:        "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png", // 默认头像
		LastLoginTime: time.Now(),
	}

	if err := database.DB.Create(&archive).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建存档失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": archive,
	})
}
