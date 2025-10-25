package config

import (
	"gin-scaffold/pkg/middleware"
)

type Middleware struct {
	Auth      middleware.AuthConfig        `mapstructure:"auth"`
	CopyBody  middleware.CopyBodyConfig    `mapstructure:"copy-body"`
	TraceId   middleware.TraceIdConfig     `mapstructure:"trace-id"`
	Cors      middleware.CorsConfig        `mapstructure:"cors"`
	Logger    middleware.LoggerConfig      `mapstructure:"logger"`
	Gzip      middleware.GzipConfig        `mapstructure:"gzip"`
	RateLimit middleware.RateLimiterConfig `mapstructure:"rate-limit"`
}
