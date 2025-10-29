package system

import (
	"context"
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/core"
	"gorm.io/gorm"
)

type roleService struct {
	repo system.RoleRepo
}

func (r *roleService) Update(ctx context.Context, role system.Role) error {
	_, err := r.repo.UpdateById(ctx, role.ID, role)
	return err
}

func (r *roleService) Query(ctx context.Context, param system.RoleQueryParam) (system.Roles, *core.Pagination, error) {
	return r.repo.FindWithPage(ctx, param.PageParam, func(db *gorm.DB) {
		db.Select("*")
	})

}

func (r *roleService) Delete(ctx context.Context, ids []uint64) error {
	return r.repo.Delete(ctx, ids)
}

func NewRoleService(repo system.RoleRepo) system.RoleService {
	return &roleService{repo: repo}
}

func (r *roleService) Create(ctx context.Context, role system.Role) error {
	return r.repo.Create(ctx, &role)
}
