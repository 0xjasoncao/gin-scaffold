package cache

import (
	"context"
	"fmt"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/redis/go-redis/v9"
	"time"
)

// Redis Redis缓存实现
type Redis struct {
	Client redis.UniversalClient
}

// NewRedis 创建单节点Redis缓存实例（带错误详情日志）
func NewRedis(ctx context.Context, addr string, password string, db int) (*Redis, func(), error) {

	// 1. 初始化客户端（明确配置信息）
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pingCmd := cli.Ping(ctx)
	if err := pingCmd.Err(); err != nil {
		logging.WithContext(ctx).Sugar().Errorf(
			"redis single node connection failed: addr=%s, db=%d, error=%v",
			addr, db, err,
		)
		_ = cli.Close()
		return nil, nil, fmt.Errorf("connect redis single node failed: %w", err)
	}

	// 5. 成功日志（可选，便于调试环境确认连接状态）
	logging.WithContext(ctx).Sugar().Infof(
		"redis single node connected successfully: addr=%s, db=%d, ping response=%s",
		addr, db, pingCmd.Val(),
	)

	cleanup := func() {
		if err := cli.Close(); err != nil {
			logging.WithContext(ctx).Sugar().Warnf(
				"redis single node close failed: addr=%s, db=%d, error=%v",
				addr, db, err,
			)
		} else {
			logging.WithContext(ctx).Sugar().Infof(
				"redis single node closed successfully: addr=%s, db=%d",
				addr, db,
			)
		}
	}

	return &Redis{Client: cli}, cleanup, nil
}

// NewClusterRedis 创建Redis集群实例（带错误详情日志）
func NewClusterRedis(ctx context.Context, addrs []string, password string) (*Redis, func(), error) {

	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	pingCmd := cli.Ping(ctx)
	if err := pingCmd.Err(); err != nil {
		logging.WithContext(ctx).Sugar().Errorf(
			"redis cluster connection failed: addrs=%v, error=%v",
			addrs, err,
		)
		_ = cli.Close()
		return nil, nil, fmt.Errorf("connect redis cluster failed: %w", err)
	}

	logging.WithContext(ctx).Sugar().Infof(
		"redis cluster connected successfully: addrs=%v, ping response=%s",
		addrs, pingCmd.Val(),
	)

	cleanup := func() {
		if err := cli.Close(); err != nil {
			logging.WithContext(ctx).Sugar().Warnf(
				"redis cluster close failed: addrs=%v, error=%v",
				addrs, err,
			)
		} else {
			logging.WithContext(ctx).Sugar().Infof(
				"redis cluster closed successfully: addrs=%v",
				addrs,
			)
		}
	}

	return &Redis{Client: cli}, cleanup, nil
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
