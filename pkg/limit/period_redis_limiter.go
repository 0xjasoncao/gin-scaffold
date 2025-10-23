package limit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisPeriodLimiter struct {
	redisClient redis.UniversalClient
	luaScript   *redis.Script
}

func NewRedisPeriodLimiter(redisClient redis.UniversalClient) *RedisPeriodLimiter {
	lua := `
	local key = KEYS[1]
	local max_count = tonumber(ARGV[1])
	local window = tonumber(ARGV[2])

	local current = redis.call("GET", key)
	if current then
    current = tonumber(current)
   		 if current >= max_count then
      		 local ttl = redis.call("PTTL", key)
             return {0, current, ttl}
   		 end
     end
	 current = redis.call("INCR", key)

	if current == 1 then
    redis.call("PEXPIRE", key, window)
	end

	local allowed = current <= max_count
	local ttl = redis.call("PTTL", key)

	return { allowed and 1 or 0, current, ttl }
`
	return &RedisPeriodLimiter{
		redisClient: redisClient,
		luaScript:   redis.NewScript(lua),
	}
}

// Allow 检查是否允许访问
// maxCount  每窗口允许次数
// window  窗口时间 单位毫秒
// return allowed,current,ttl,error
func (r *RedisPeriodLimiter) Allow(ctx context.Context, key string, maxCount int, windowMs int64) (bool, int64, int64, error) {

	res, err := r.luaScript.Run(ctx, r.redisClient, []string{key}, maxCount, windowMs).Result()
	if err != nil {
		return false, 0, 0, err
	}
	arr, ok := res.([]interface{})
	if !ok || len(arr) < 3 {
		return false, 0, 0, nil
	}
	allowed := arr[0].(int64) == 1
	count := arr[1].(int64)
	ttl := arr[2].(int64)
	return allowed, count, ttl, nil
}
