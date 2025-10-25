package limiter

import "context"

type PeriodLimiter interface {
	Allow(ctx context.Context, key string, maxCount int, windowMs int64) (bool, int64, int64, error)
}
