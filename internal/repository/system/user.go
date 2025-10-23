package system

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/repo"
	"gorm.io/gorm"
)

type userRepo struct {
	repo.BasicRepo[system.User]
}

func NewUserRepo(db *gorm.DB) system.UserRepo {
	return &userRepo{repo.NewBasicRepo[system.User](db)}
}
