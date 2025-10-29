package system

import (
	"context"

	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/core/errorsx"
	"gin-scaffold/pkg/utils/encryptutil"
	"gorm.io/gorm"
)

type userService struct {
	userRepo system.UserRepo
}

func NewUserService(userRepo system.UserRepo) system.UserService {
	return &userService{userRepo: userRepo}
}

func (srv *userService) Login(ctx context.Context, param system.UserQueryParam) (*system.User, error) {
	if param.Email != "" {

		u, err := srv.userRepo.FindByWhere(ctx, "email=?", param.Email)
		if err != nil {
			if errorsx.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorsx.NewNotFound("用户不存在")
			}
			return nil, err
		}
		if !encryptutil.VerifyPassword(u.PasswordHash, param.Password) {
			return nil, errorsx.NewUnauthorized("密码输入错误")
		}
		if u.IsDisable() {
			return nil, errorsx.NewForbidden("用户被禁用,请联系管理员!")
		}
		return u, nil
	}
	return nil, errorsx.NewBadRequest("参数错误")
}
