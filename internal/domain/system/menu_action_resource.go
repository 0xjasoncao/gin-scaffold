package system

import "gin-scaffold/internal/domain/shared"

type Model struct {
	shared.BasicInfo
	ActionID string `gorm:"column:action_id;size:36;index;default:'';not null;"`
	Method   string `gorm:"column:method;size:100;default:'';not null;"`
	Path     string `gorm:"column:path;size:100;default:'';not null;"`
}

func (Model) TableName() string {
	return "sys_menu_action_resources"
}
