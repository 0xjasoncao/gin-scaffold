package user

import (
	"time"
)

// User 持久化对象（对应数据库表结构）
type User struct {
	ID           string     `gorm:"primaryKey;type:varchar(36);comment:用户唯一标识ID"`
	Name         string     `gorm:"size:100;not null;comment:用户真实姓名"`
	PasswordHash string     `gorm:"column:password;type:varchar(255);not null;comment:密码哈希值"`
	NickName     string     `gorm:"size:100;comment:用户昵称"`
	Email        string     `gorm:"size:255;index;comment:电子邮箱"`
	Status       int        `gorm:"type:tinyint;default:0;comment:用户状态（0-未激活，1-已激活，2-已封禁）"`
	Mobile       string     `gorm:"size:20;index;comment:手机号码"`
	Gender       int        `gorm:"type:tinyint;comment:性别（0-未知，1-男，2-女）"`
	Birthday     *time.Time `gorm:"type:date;comment:出生日期"`
	Introduction string     `gorm:"type:text;comment:个人简介"`
	CreatedAt    *time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime;comment:更新时间"`
	CreatedBy    string     `gorm:"size:36;comment:创建人ID"`
	UpdatedBy    string     `gorm:"size:36;comment:更新人ID"`
}

func (User) TableName() string {
	return "user"
}
