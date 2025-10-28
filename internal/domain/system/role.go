package system

import (
	"context"
	"gin-scaffold/internal/domain/shared"
	BasicRepo "gin-scaffold/pkg/repo"
)

type Role struct {
	shared.BasicInfo
	Name    string `gorm:"column:name;size:100;index;default:'';not null;"`
	Comment string `gorm:"column:comment;size:1024;"`
	Status  int    `gorm:"column:status;index;default:0;not null;"`
}

type Roles []*Role

func (Role) TableName() string {
	return "sys_roles"
}

type RoleRepo interface {
	BasicRepo.IRepo[Role]
}

type RoleService interface {
	Create(ctx context.Context, role Role) error
	Delete(ctx context.Context, ids []uint64) error
}
