package user

import (
	"github.com/0xjasoncao/gin-scaffold/internal/model"
	"time"
)

// User 持久化对象（对应数据库表结构）
type User struct {
	ID           string    `gorm:"primaryKey;type:varchar(36)"`
	Name         string    `gorm:"size:100;not null"`
	Account      string    `gorm:"size:50;uniqueIndex;not null"`
	PasswordHash string    `gorm:"column:password;type:varchar(255);not null"`
	NickName     string    `gorm:"size:100"`
	Email        string    `gorm:"size:255;index"`
	Status       int       `gorm:"type:tinyint;default:0"`
	Mobile       string    `gorm:"size:20;index"`
	Gender       int       `gorm:"type:tinyint"`
	Birthday     time.Time `gorm:"type:date"`
	Introduction string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	CreatedBy    string    `gorm:"size:36"`
	UpdatedBy    string    `gorm:"size:36"`
}

func (User) TableName() string {
	return "user"
}

func init() {
	model.Register(&User{})
}
