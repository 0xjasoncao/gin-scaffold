package config

import (
	middleware2 "gin-scaffold/pkg/core/middleware"
)

type Middleware struct {
	Auth      middleware2.AuthConfig        `mapstructure:"auth"`
	CopyBody  middleware2.CopyBodyConfig    `mapstructure:"copy-body"`
	TraceId   middleware2.TraceIdConfig     `mapstructure:"trace-id"`
	Cors      middleware2.CorsConfig        `mapstructure:"cors"`
	Logger    middleware2.LoggerConfig      `mapstructure:"logger"`
	Gzip      middleware2.GzipConfig        `mapstructure:"gzip"`
	RateLimit middleware2.RateLimiterConfig `mapstructure:"rate-limit"`
}
