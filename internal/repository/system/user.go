package system

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/core"
	"gorm.io/gorm"
)

type userRepo struct {
	core.Repository[system.User]
}

func NewUserRepo(db *gorm.DB) system.UserRepo {
	return &userRepo{core.NewRepository[system.User](db)}
}
