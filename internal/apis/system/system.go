package system

import (
	"gin-scaffold/internal/apis/system/v1"
	"github.com/google/wire"
)

type V1 struct {
	User *v1.UserHandler
}

type Handlers struct {
	V1 *V1
}

var ProviderSet = wire.NewSet(
	v1.NewUserHandler,
	wire.Struct(new(V1), "*"),
	wire.Struct(new(Handlers), "*"),
)
