package repository

import (
	"github.com/0xjasoncao/gin-scaffold/internal/repository/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.NewUserRepo,
)
