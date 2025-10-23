package system

import (
	"context"
	"gin-scaffold/pkg/repo"
	"time"

	"gin-scaffold/internal/domain/shared"
)

// User 持久化对象（对应数据库表结构）
type User struct {
	shared.BasicInfo
	Name         string     `gorm:"size:100;not null"`
	PasswordHash string     `gorm:"column:password;type:varchar(255);not null"`
	NickName     string     `gorm:"size:100"`
	Email        string     `gorm:"size:255;index"`
	HeaderImage  string     `gorm:"size:300"`
	Mobile       string     `gorm:"size:20;index"`
	Gender       int        `gorm:"type:tinyint"`
	Birthday     *time.Time `gorm:"type:date"`
	Disable      int        `gorm:"type:tinyint;default:0"`
	Introduction string     `gorm:"type:text"`
}

type UserQueryParam struct {
	Password   string
	VerifyCode string
	Mobile     string
	Email      string
}

func (User) TableName() string {
	return "sys_users"
}

type UserRepo interface {
	repo.IRepo[User]
}

type UserService interface {
	Login(ctx context.Context, param UserQueryParam) (*User, error)
}

func (u User) IsDisable() bool {
	return u.Disable == 1
}
