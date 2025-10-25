package service

import (
	"gin-scaffold/internal/service/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewUserService,
	system.NewRoleService,
)
