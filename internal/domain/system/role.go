package system

import (
	"context"
	"gin-scaffold/internal/domain/shared"
	core "gin-scaffold/pkg/core"
)

type Role struct {
	shared.BasicInfo
	Name    string `gorm:"column:name;size:100;index;default:'';not null;"`
	Comment string `gorm:"column:comment;size:1024;"`
	Status  int    `gorm:"column:status;index;default:0;not null;"`
}

type Roles []*Role

type RoleQueryParam struct {
	core.PageParam
}

func (Role) TableName() string {
	return "sys_roles"
}

type RoleRepo interface {
	core.UniversalRepo[Role]
}

type RoleService interface {
	Create(ctx context.Context, role Role) error
	Delete(ctx context.Context, ids []uint64) error
	Update(ctx context.Context, role Role) error
	Query(ctx context.Context, param RoleQueryParam) (Roles, *core.Pagination, error)
}
