package handler

import (
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler/V1/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.NewHandler,
	user.NewLoginHandler,
	wire.Struct(new(V1), "*"),
	wire.Struct(new(Handler), "*"),
)
