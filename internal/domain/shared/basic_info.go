package shared

import (
	"gin-scaffold/pkg/core/ginutil"
	"gin-scaffold/pkg/sonyflakex"
	"gorm.io/gorm"
	"time"
)

type BasicInfo struct {
	ID        uint64 `gorm:"primaryKey;"`
	CreatorId uint64 `gorm:"column:creator_id;size:40;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;"`
}

// BeforeCreate gorm before create hook
func (b *BasicInfo) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = sonyflakex.NewSonyFlakeId()
	tokenInfo := ginutil.TokenFromContext(tx.Statement.Context)
	if tokenInfo != nil {
		b.CreatorId = tokenInfo.UserID
	}
	return
}
