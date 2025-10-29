package system

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/core"
	"gorm.io/gorm"
)

type roleRepo struct {
	core.Repository[system.Role]
}

func NewRoleRepo(db *gorm.DB) system.RoleRepo {
	return &roleRepo{core.NewRepository[system.Role](db)}
}
