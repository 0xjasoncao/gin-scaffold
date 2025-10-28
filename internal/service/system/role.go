package system

import (
	"context"
	"gin-scaffold/internal/domain/system"
)

type roleService struct {
	repo system.RoleRepo
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
