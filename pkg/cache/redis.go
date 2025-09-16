package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Redis Redis缓存实现
type Redis struct {
	Client redis.UniversalClient
}

// NewRedis 创建新的Redis缓存实例
func NewRedis(addr string, password string, db int) (*Redis, func()) {
	Client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Redis{
			Client: Client,
		}, func() {
			Client.Close()
		}
}

func NewClusterRedis(addrs []string, password string) (*Redis, func()) {
	Client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})
	return &Redis{
			Client: Client,
		}, func() {
			Client.Close()
		}
}

// Set 存储键值对
func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

// Get 获取键对应的值
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Delete 删除指定键
func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Exists 判断键是否存在
func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Expire 设置键的过期时间
func (r *Redis) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.Client.Expire(ctx, key, expiration).Err()
}

// Incr 对键的值进行自增
func (r *Redis) Incr(ctx context.Context, key string) (int64, error) {
	return r.Client.Incr(ctx, key).Result()
}

// Decr 对键的值进行自减
func (r *Redis) Decr(ctx context.Context, key string) (int64, error) {
	return r.Client.Decr(ctx, key).Result()
}

// Close 关闭缓存连接
func (r *Redis) Close() error {
	return r.Client.Close()
}
