package provider

import (
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
	"github.com/0xjasoncao/gin-scaffold/pkg/token"
)

func InitTokenService(config *config.Config,
	redis *cache.Redis,
	mem *cache.Memory,
) (token.Service, func(), error) {
	cfg := config.Jwt

	var c cache.Cache
	if cfg.Store == "redis" {
		c = redis
	} else {
		c = mem
	}
	store := token.NewStoreWithCache(c)

	service, err := token.NewTokenService(&token.Settings{
		ExpiresAtSeconds: cfg.ExpiresAtSeconds,
		Issuer:           cfg.Issuer,
		Key:              cfg.Key,
	}, store)
	if err != nil {
		return nil, func() {}, err
	}
	return service, func() {}, nil

}
