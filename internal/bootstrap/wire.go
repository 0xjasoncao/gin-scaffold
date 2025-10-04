//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"

	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler"
	"github.com/0xjasoncao/gin-scaffold/internal/repository"
	"github.com/0xjasoncao/gin-scaffold/internal/service"

	"github.com/0xjasoncao/gin-scaffold/internal/apis/router"
	"github.com/0xjasoncao/gin-scaffold/internal/bootstrap/provider"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	provider.BasicProviderSet,
	router.NewRouter,
	repository.ProviderSet,
	service.ProviderSet,
	handler.ProviderSet,
	wire.Struct(new(ApiInjector), "*"),
)

func BuildInjector(ctx context.Context, config *config.Config) (*ApiInjector, func(), error) {
	wire.Build(ProviderSet)
	return nil, nil, nil
}
