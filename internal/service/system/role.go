package system

import (
	"context"
	"gin-scaffold/internal/domain/system"
)

type roleService struct {
	repo system.RoleRepo
}

func NewRoleService(repo system.RoleRepo) system.RoleService {
	return &roleService{repo: repo}
}

func (r *roleService) Create(ctx context.Context, role system.Role) error {
	return r.repo.Create(ctx, &role)
}
