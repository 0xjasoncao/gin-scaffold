package system

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/repo"
	"gorm.io/gorm"
)

type roleRepo struct {
	repo.BasicRepo[system.Role]
}

func NewRoleRepo(db *gorm.DB) system.RoleRepo {
	return &roleRepo{repo.NewBasicRepo[system.Role](db)}
}
