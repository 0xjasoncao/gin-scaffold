package middleware

import (
	"fmt"
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/limit"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/utils/encryptutil"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type RateLimiterConfig struct {
	Enable            bool     `mapstructure:"enable"`
	SkipPathPrefix    []string `mapstructure:"skip-path-prefix"`
	NotSkipPathPrefix []string `mapstructure:"not-skip-path-prefix"`
	Rules             []struct {
		URLPattern string         `mapstructure:"url-pattern"` //请求url正则匹配
		Period     int64          `mapstructure:"period"`      //窗口时间，单位:秒
		Strategy   string         `mapstructure:"strategy"`    //限流策略(ip|user)
		MaxCount   int            `mapstructure:"max-count"`   //最大请求次数
		Regex      *regexp.Regexp `mapstructure:"-"`
	} `mapstructure:"rules"`
}

const (
	IPStrategy       = "ip"
	UserStrategy     = "user"
	RateLimitUserKey = "rate_limit:user:%d:%s_%d"
	RateLimitIpKey   = "rate_limit:ip:%s:%s_%d"
)

func RateLimitMiddleware(redisClient redis.UniversalClient, conf *RateLimiterConfig) gin.HandlerFunc {
	limiter := limit.NewRedisPeriodLimiter(redisClient)

	//预编译正则 避免每次请求编译
	for i := range conf.Rules {
		if conf.Rules[i].URLPattern != "" {
			conf.Rules[i].Regex = regexp.MustCompile(conf.Rules[i].URLPattern)
		}
	}

	return func(c *gin.Context) {
		if NeedSkip(c, SkippedPathPrefix(conf.SkipPathPrefix...), NotSkippedPathPrefix(conf.NotSkipPathPrefix...)) {
			c.Next()
			return
		}
		ctx := c.Request.Context()
		tokenInfo := api.TokenFromContext(ctx)
		requestIP := c.ClientIP()
		requestURL := c.Request.URL.Path

		for i, rule := range conf.Rules {
			if rule.Regex != nil && rule.Regex.Match([]byte(requestURL)) {
				var (
					allowed bool
					err     error
				)
				var redisKey string
				switch strings.ToLower(rule.Strategy) {
				case UserStrategy:
					if tokenInfo == nil {
						continue
					}
					redisKey = fmt.Sprintf(RateLimitUserKey, tokenInfo.UserID, encryptutil.Md5(rule.URLPattern), i)
				case IPStrategy:
					redisKey = fmt.Sprintf(RateLimitIpKey, requestIP, encryptutil.Md5(rule.URLPattern), i)
				}
				allowed, _, _, err = limiter.Allow(c.Request.Context(), redisKey, rule.MaxCount, rule.Period*1000)

				if err != nil {
					logging.WithContext(ctx).Error("Rate limiter middleware error", zap.Error(err))
					api.ResError(c, errorsx.NewInternal("Internal server error, please try again later."))
					return
				}
				if !allowed {
					logging.WithContext(ctx).Warn("Rate limit triggered",
						zap.String("url", requestURL),
						zap.String("strategy", rule.Strategy),
						zap.Int("max_count", rule.MaxCount),
						zap.Int64("period_ms", rule.Period),
						zap.String("url_pattern", rule.URLPattern),
						zap.String("redis_key", redisKey),
					)
					api.ResError(c, errorsx.NewTooManyRequest("Too many requests, please try again later."))
					return
				}
			}

		}
		c.Next()

	}
}
