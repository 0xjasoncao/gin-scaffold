package redisx

import (
	"context"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/logging"
	"github.com/redis/go-redis/v9"
	"slices"
)

type (
	Factory struct {
		clients map[int]redis.UniversalClient
		conf    *Config
	}
	Config struct {
		//多个地址为集群模式
		Addrs []string `mapstructure:"addrs"`
		//仅在单机模式下有用
		DBs      []int  `mapstructure:"dbs"`
		Password string `mapstructure:"password"`
	}
)

func (c *Config) isCluster() bool {
	return c != nil && len(c.Addrs) > 1
}

// NewRedisFactory 创建Redis客户端工厂
// 单机模式：根据提供的db切片，注册不同db的客户端实例，默认会注册db0的客户端实例
// 集群模式：只注册db0的客户端实例
func NewRedisFactory(ctx context.Context, conf *Config) (*Factory, func(), error) {
	factory := &Factory{
		clients: make(map[int]redis.UniversalClient),
		conf:    conf,
	}
	if conf.Addrs == nil || len(conf.Addrs) == 0 {
		return nil, func() {}, errorsx.New("Redis addresses can not be empty")
	}

	// register redis clients
	if !conf.isCluster() {
		//single node
		dbs := conf.DBs
		if !slices.Contains(dbs, 0) {
			dbs = append(dbs, 0)
		}
		addr := conf.Addrs[0]
		var funcs []func()
		for _, db := range conf.DBs {
			cli, f, err := NewSingle(ctx, addr, db, conf.Password)
			if err != nil {
				return nil, f, err
			}
			factory.clients[db] = cli
			funcs = append(funcs, f)
		}
		return factory, func() {
			for _, f := range funcs {
				f()
			}
		}, nil
	} else {
		//cluster
		cluster, f, err := NewCluster(ctx, conf.Addrs, conf.Password)
		if err != nil {
			return nil, f, err
		}
		factory.clients[0] = cluster
		return factory, f, nil
	}
}

// Get 根据db获取对应的Redis实例
func (f *Factory) Get(db int) (redis.UniversalClient, error) {
	if f.conf.isCluster() {
		return f.clients[0], nil
	}
	client, ok := f.clients[db]
	if !ok {
		return nil, errorsx.Errorf("Redis db: %d not exists", db)
	}
	return client, nil
}

// GetDefault 默认获取db0的客户端实例
func (f *Factory) GetDefault() redis.UniversalClient {
	return f.clients[0]
}

// NewSingle 创建单节点Redis缓存实例
func NewSingle(ctx context.Context, addr string, db int, password string) (redis.UniversalClient, func(), error) {

	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pingCmd := cli.Ping(ctx)
	if err := pingCmd.Err(); err != nil {
		logging.WithContext(ctx).Sugar().Errorf(
			"Redis single node connection failed: addr=%s, db=%d, error=%v",
			addr, db, err,
		)
		_ = cli.Close()
		return nil, nil, errorsx.Errorf("Connect redis single node failed: %w", err)
	}

	logging.WithContext(ctx).Sugar().Infof(
		"Redis single node connected successfully: addr=%s, db=%d, ping response=%s",
		addr, db, pingCmd.Val(),
	)

	cleanup := func() {
		if err := cli.Close(); err != nil {
			logging.WithContext(ctx).Sugar().Warnf(
				"Redis single node close failed: addr=%s, db=%d, error=%v",
				addr, db, err,
			)
		} else {
			logging.WithContext(ctx).Sugar().Infof(
				"Redis single node closed successfully: addr=%s, db=%d",
				addr, db,
			)
		}
	}
	return cli, cleanup, nil
}

// NewCluster 创建Redis集群实例
func NewCluster(ctx context.Context, addrs []string, password string) (redis.UniversalClient, func(), error) {
	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	pingCmd := cli.Ping(ctx)
	if err := pingCmd.Err(); err != nil {
		logging.WithContext(ctx).Sugar().Errorf(
			"Redis cluster connection failed: addrs=%v, error=%v",
			addrs, err,
		)
		_ = cli.Close()
		return nil, nil, errorsx.Errorf("Connect redis cluster failed: %w", err)
	}

	logging.WithContext(ctx).Sugar().Infof(
		"Redis cluster connected successfully: addrs=%v, ping response=%s",
		addrs, pingCmd.Val(),
	)

	cleanup := func() {
		if err := cli.Close(); err != nil {
			logging.WithContext(ctx).Sugar().Warnf(
				"Redis cluster close failed: addrs=%v, error=%v",
				addrs, err,
			)
		} else {
			logging.WithContext(ctx).Sugar().Infof(
				"Redis cluster closed successfully: addrs=%v",
				addrs,
			)
		}
	}

	return cli, cleanup, nil
}
