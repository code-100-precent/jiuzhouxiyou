package handlers

import (
	"net/http"

	"lingecho-backend/database"
	"lingecho-backend/models"
	"lingecho-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UID string `json:"uid" binding:"required"`
	Pwd string `json:"pwd" binding:"required"`
}

type RegisterRequest struct {
	UID    string `json:"uid" binding:"required,min=3,max=20"`
	Pwd    string `json:"pwd" binding:"required,min=6"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误: " + err.Error(),
		})
		return
	}

	// 查找用户
	var user models.User
	if err := database.DB.Where("uid = ?", req.UID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "账号或密码错误",
		})
		return
	}

	// 验证密码
	if !utils.CheckPassword(req.Pwd, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "账号或密码错误",
		})
		return
	}

	// 生成 token
	token, err := utils.GenerateToken(user.ID, user.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成 token 失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": LoginResponse{
			Token: token,
		},
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误: " + err.Error(),
		})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := database.DB.Where("uid = ?", req.UID).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": 409,
			"msg":  "用户名已存在",
		})
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Pwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "密码加密失败",
		})
		return
	}

	// 设置默认值
	if req.Name == "" {
		req.Name = req.UID
	}
	if req.Avatar == "" {
		req.Avatar = "https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png"
	}

	// 创建用户
	user := models.User{
		UID:      req.UID,
		Password: hashedPassword,
		Name:     req.Name,
		Avatar:   req.Avatar,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "注册失败: " + err.Error(),
		})
		return
	}

	// 生成 token
	token, err := utils.GenerateToken(user.ID, user.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成 token 失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": LoginResponse{
			Token: token,
		},
	})
}

// Ping 心跳检测
func Ping(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := c.Get("uid")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "pong",
		"data": gin.H{
			"user_id": userID,
			"uid":     uid,
		},
	})
}
