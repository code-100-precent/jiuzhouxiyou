package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UID       string         `gorm:"uniqueIndex;not null;size:100" json:"uid"`
	Password  string         `gorm:"not null;size:255" json:"-"` // 不返回密码
	Name      string         `gorm:"size:100" json:"name"`
	Avatar    string         `gorm:"size:500" json:"avatar"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Archives []Archive `gorm:"foreignKey:UserID" json:"archives,omitempty"`
}

// Archive 存档表
type Archive struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"not null;index" json:"user_id"`
	Name          string         `gorm:"size:100" json:"name"`
	Job           string         `gorm:"size:50" json:"job"`
	Avatar        string         `gorm:"size:500" json:"avatar"`
	LastLoginTime time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"last_login_time"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	User User `gorm:"foreignKey:UserID" json:"-"`
}
