package provider

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
)

func InitCache(
	config *config.Config,
	redis *cache.Redis,
	mem *cache.Memory) cache.Cache {

	switch config.Cache.Use {
	case "redis":
		return redis
	default:
		return mem
	}

}

func InitRedisCli(ctx context.Context, config *config.Config) (*cache.Redis, func(), error) {

	redisConf := config.Redis
	if !redisConf.Open {
		logging.WithContext(ctx).Sugar().Warnf("Redis is not enabled (configured to closed), skip initialization")
		return nil, func() {}, nil
	}
	if !redisConf.UseCluster {
		redis, cleanFunc, err := cache.NewRedis(ctx, redisConf.Addr, redisConf.Password, redisConf.DB)
		return redis, cleanFunc, err
	} else {
		redis, cleanFunc, err := cache.NewClusterRedis(ctx, redisConf.ClusterAddress, redisConf.Password)
		return redis, cleanFunc, err
	}
}

func InitMemoryCache() (*cache.Memory, func()) {
	return cache.NewMemory()
}
