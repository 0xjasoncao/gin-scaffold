//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler/V1/auth"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler/V1/user"
	"github.com/0xjasoncao/gin-scaffold/internal/service"

	"github.com/0xjasoncao/gin-scaffold/internal/apis/router"
	"github.com/0xjasoncao/gin-scaffold/internal/bootstrap/provider"
	"github.com/google/wire"
)

var RepoProviderSet = wire.NewSet()

var HandlerProviderSet = wire.NewSet(
	user.NewHandler,
	auth.NewHandler,
	wire.Struct(new(handler.V1), "*"),
	wire.Struct(new(handler.Handler), "*"),
)

var ProviderSet = wire.NewSet(
	provider.BasicProviderSet,
	router.NewRouter,
	RepoProviderSet,
	service.ProviderSet,
	HandlerProviderSet,
	wire.Struct(new(ApiInjector), "*"),
)

func BuildInjector(ctx context.Context, config *config.Config) (*ApiInjector, func(), error) {
	wire.Build(ProviderSet)
	return nil, nil, nil
}
