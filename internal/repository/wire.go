package repository

import (
	"gin-scaffold/internal/repository/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewUserRepo,
	system.NewRoleRepo,
)
