package token

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
	"time"
)

type Store interface {
	Set(ctx context.Context, tokenStr string, expiration time.Duration) error
	Delete(ctx context.Context, tokenStr string) error
	Check(ctx context.Context, tokenStr string) (bool, error)
}

type store struct {
	c cache.Cache
}

func NewStoreWithCache(c cache.Cache) Store {
	return &store{c: c}
}

func (s *store) Set(ctx context.Context, tokenStr string, expiration time.Duration) error {
	return s.c.Set(ctx, tokenStr, "1", expiration)
}

func (s *store) Check(ctx context.Context, tokenStr string) (bool, error) {
	return s.c.Exists(ctx, tokenStr)
}
func (s *store) Delete(ctx context.Context, tokenStr string) error {
	return s.c.Delete(ctx, tokenStr)
}
