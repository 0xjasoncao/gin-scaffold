package service

import (
	"context"

	"github.com/0xjasoncao/gin-scaffold/internal/domain/user"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/encrypt"
	"gorm.io/gorm"
)

type LoginService interface {
	Login(ctx context.Context, opt LoginOpt) (*user.User, error)
	Logout(ctx context.Context, accessToken string) error
}

type loginService struct {
	userRepo user.Repo
}

func NewLoginService(userRepo user.Repo) LoginService {
	return &loginService{userRepo: userRepo}
}

type LoginOpt struct {
	Password   string
	VerifyCode string
	Mobile     string
}

func (srv *loginService) Login(ctx context.Context, opt LoginOpt) (*user.User, error) {
	if opt.Mobile != "" {
		user, err := srv.userRepo.FindUserByMobile(ctx, opt.Mobile)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.NewNotFound("用户不存在")
			}
			return nil, err
		}
		if !encrypt.VerifyPassword(user.PasswordHash, opt.Password) {
			return nil, errors.NewUnauthorized("密码输入错误")
		}
		return user, nil

	}

	return nil, errors.NewBadRequest("参数错误")
}

func (srv *loginService) Logout(ctx context.Context, accessToken string) error {
	return nil
}
