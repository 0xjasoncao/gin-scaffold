package system

import "gin-scaffold/internal/domain/shared"

type RoleMenu struct {
	shared.BasicInfo
	RoleID   string `gorm:"column:role_id;size:36;index;default:'';not null;"`
	MenuID   string `gorm:"column:menu_id;size:36;index;default:'';not null;"`
	ActionID string `gorm:"column:action_id;size:36;index;default:'';not null;"`
}

type RoleMenus []*RoleMenu

func (RoleMenu) TableName() string {
	return "sys_role_menus"
}
