package provider

import (
	"context"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/redisx"

	"gin-scaffold/internal/config"
)

func InitRedisFactory(ctx context.Context, config *config.Config) (*redisx.Factory, func(), error) {
	logging.WithContext(ctx).Sugar().Infof("[Redis] - Initializing redis factory...")
	return redisx.NewRedisFactory(ctx, &config.Redis)
}
