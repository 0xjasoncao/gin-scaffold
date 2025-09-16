package cache

import (
	"context"
	"time"
)

// Cache 缓存接口定义
type Cache interface {
	// Set 存储键值对
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Get 获取键对应的值，将结果存入value指针
	Get(ctx context.Context, key string) (string, error)

	// Delete 删除指定键
	Delete(ctx context.Context, key string) error

	// Exists 判断键是否存在
	Exists(ctx context.Context, key string) (bool, error)

	// Expire 设置键的过期时间
	Expire(ctx context.Context, key string, expiration time.Duration) error

	// Incr 对键的值进行自增
	Incr(ctx context.Context, key string) (int64, error)

	// Decr 对键的值进行自减
	Decr(ctx context.Context, key string) (int64, error)

	// Close 关闭缓存连接
	Close() error
}
