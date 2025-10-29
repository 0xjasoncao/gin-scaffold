package apis

import (
	"gin-scaffold/internal/apis/swagger"
	"gin-scaffold/internal/apis/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	swagger.NewHandler,
	system.ProviderSet,
	wire.Struct(new(RouterHandlers), "*"),
)
