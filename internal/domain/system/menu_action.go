package system

import "gin-scaffold/internal/domain/shared"

type MenuAction struct {
	shared.BasicInfo
	MenuID string `gorm:"column:menu_id;size:36;index;default:'';not null;"`
	Code   string `gorm:"column:code;size:100;default:'';not null;"`
	Name   string `gorm:"column:name;size:100;default:'';not null;"`
}

func (MenuAction) TableName() string {
	return "sys_menu_actions"
}
