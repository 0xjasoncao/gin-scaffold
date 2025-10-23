//go:build wireinject
// +build wireinject

package initialize

import (
	"context"
	"gin-scaffold/internal/apis"
	"gin-scaffold/internal/config"
	"github.com/gin-gonic/gin"

	"gin-scaffold/internal/repository"
	"gin-scaffold/internal/service"

	"gin-scaffold/internal/initialize/provider"
	"github.com/google/wire"
)

type ApiInjector struct {
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	provider.BasicProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
	apis.ProviderSet,
	wire.Struct(new(ApiInjector), "*"),
)

func BuildInjector(ctx context.Context, config *config.Config) (*ApiInjector, func(), error) {
	wire.Build(ProviderSet)
	return nil, nil, nil
}
