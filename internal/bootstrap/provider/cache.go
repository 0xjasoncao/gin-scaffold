package provider

import (
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
)

func InitCache(
	config *config.Config,
	redis *cache.Redis,
	mem *cache.Memory) cache.Cache {
	if config.Cache.Use == "redis" {
		return redis
	}
	return mem

}

func InitRedisCli(config *config.Config) (*cache.Redis, func()) {
	redisConf := config.Redis
	if !redisConf.UseCluster {
		Redis, cleanFunc := cache.NewRedis(redisConf.Addr, redisConf.Password, redisConf.DB)
		return Redis, cleanFunc
	} else {
		Redis, cleanFunc := cache.NewClusterRedis(redisConf.ClusterAddress, redisConf.Password)
		return Redis, cleanFunc
	}
}

func InitMemoryCache() (*cache.Memory, func()) {
	return cache.NewMemory()
}
