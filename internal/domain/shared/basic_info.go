package shared

import (
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/sonyflakex"
	"gorm.io/gorm"
	"time"
)

type BasicInfo struct {
	ID        uint64 `gorm:"primaryKey;"`
	Creator   uint64 `gorm:"column:creator;size:10;"`
	CreatorId string `gorm:"column:creator_id;size:40;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;"`
}

// BeforeCreate gorm before create hook
func (b *BasicInfo) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = sonyflakex.NewSonyFlakeId()
	tokenInfo := api.TokenFromContext(tx.Statement.Context)
	if tokenInfo != nil {
		b.Creator = tokenInfo.UserID
	}
	return
}
