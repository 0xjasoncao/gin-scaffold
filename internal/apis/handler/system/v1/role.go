package v1

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/token"
)

type RoleHandler struct {
	UserSrv  system.UserService
	TokenSrv token.Service
}
