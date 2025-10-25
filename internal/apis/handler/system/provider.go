package system

import (
	v1 "gin-scaffold/internal/apis/handler/system/v1"
	"github.com/google/wire"
)

// ProviderSet wire provider
var ProviderSet = wire.NewSet(
	v1.NewUserHandler,
	wire.Struct(new(V1), "*"),
	wire.Struct(new(Handlers), "*"),
)
