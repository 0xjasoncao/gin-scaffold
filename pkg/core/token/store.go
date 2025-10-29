package token

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"time"
)

type Store interface {
	Set(ctx context.Context, tokenStr string, expiration time.Duration) error
	Delete(ctx context.Context, tokenStr string) error
	Check(ctx context.Context, tokenStr string) (bool, error)
}

type redisStore struct {
	cli redis.UniversalClient
}

func NewRedisStore(client redis.UniversalClient) Store {
	return &redisStore{cli: client}
}

func (s *redisStore) Set(ctx context.Context, tokenStr string, expiration time.Duration) error {
	_, err := s.cli.Set(ctx, tokenStr, "1", expiration).Result()
	return err
}

func (s *redisStore) Check(ctx context.Context, tokenStr string) (bool, error) {
	result, err := s.cli.Exists(ctx, tokenStr).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}
func (s *redisStore) Delete(ctx context.Context, tokenStr string) error {
	_, err := s.cli.Del(ctx, tokenStr).Result()
	return err
}

type memoryStore struct {
	mem *cache.Cache
}

func NewMemoryStore(client *redis.Client) Store {
	return &memoryStore{mem: cache.New(0, 5*time.Minute)}
}

func (s *memoryStore) Set(ctx context.Context, tokenStr string, expiration time.Duration) error {
	s.mem.Set(tokenStr, "1", expiration)
	return nil
}

func (s *memoryStore) Check(ctx context.Context, tokenStr string) (bool, error) {
	_, exists := s.mem.Get(tokenStr)
	return exists, nil
}
func (s *memoryStore) Delete(ctx context.Context, tokenStr string) error {
	s.mem.Delete(tokenStr)
	return nil
}
