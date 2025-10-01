package service

import (
	"github.com/0xjasoncao/gin-scaffold/internal/service/auth"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	auth.NewService,
)
