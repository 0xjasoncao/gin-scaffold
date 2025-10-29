package provider

import (
	"context"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/core/token"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/redisx"
)

func InitTokenService(ctx context.Context, config *config.Config, redisFactory *redisx.Factory) (token.Service, func(), error) {
	cfg := config.Jwt
	logging.WithContext(ctx).Sugar().Infof("[JWT] - Initializing token service...")
	//默认使用db0存储
	store := token.NewRedisStore(redisFactory.GetDefault())
	service, err := token.NewTokenService(&token.Settings{
		ExpiresAtSeconds: cfg.ExpiresAtSeconds,
		Issuer:           cfg.Issuer,
		Key:              cfg.Key,
	}, store)
	return service, func() {
	}, err

}
