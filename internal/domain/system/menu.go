package system

import (
	"gin-scaffold/internal/domain/shared"
)

type Menu struct {
	shared.BasicInfo
	Name       string `gorm:"column:name;size:50;index;default:'';not null;"`
	Sequence   int    `gorm:"column:sequence;index;default:0;not null;"`
	Icon       string `gorm:"column:icon;size:255;"`
	Router     string `gorm:"column:router;size:255;"`
	ParentID   string `gorm:"column:parent_id;size:36;index;"`
	ParentPath string `gorm:"column:parent_path;size:518;index;"`
	ShowStatus int    `gorm:"column:show_status;index;default:0;not null;"`
	Status     int    `gorm:"column:status;index;default:0;not null;"`
	Comment    string `gorm:"column:comment;size:1024;"`
}

func (Menu) TableName() string {
	return "sys_menus"
}
