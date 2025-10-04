package user

import (
	"context"

	"github.com/0xjasoncao/gin-scaffold/internal/domain/user"
	BasicRepo "github.com/0xjasoncao/gin-scaffold/pkg/repo"
	"gorm.io/gorm"
)

type repo struct {
	BasicRepo.BasicRepo[user.User]
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &repo{BasicRepo.NewBasicRepo[user.User](db)}
}

func (r *repo) FindUserByMobile(ctx context.Context, mobile string) (*user.User, error) {
	return r.BasicRepo.FindByWhere(ctx, "mobile=?", mobile)

}
