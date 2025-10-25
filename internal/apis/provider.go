package apis

import (
	"gin-scaffold/internal/apis/handler/swagger"
	"gin-scaffold/internal/apis/handler/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	swagger.NewHandler,
	system.ProviderSet,
	wire.Struct(new(RouterHandlers), "*"),
)
