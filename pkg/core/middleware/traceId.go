package middleware

import (
	"strconv"

	"gin-scaffold/pkg/logging"

	"gin-scaffold/pkg/sonyflakex"
	"github.com/gin-gonic/gin"
)

const TraceId = "X-Request-Id"

type TraceIdConfig struct {
	MaxContentLength  int64    `mapstructure:"max-content-length"`
	Enable            bool     `mapstructure:"enable"`
	SkipPathPrefix    []string `mapstructure:"skip-path-prefix"`
	NotSkipPathPrefix []string `mapstructure:"not-skip-path-prefix"`
}

func Trace(config *TraceIdConfig) gin.HandlerFunc {

	return func(c *gin.Context) {
		if NeedSkip(c, SkippedPathPrefix(config.SkipPathPrefix...), NotSkippedPathPrefix(config.NotSkipPathPrefix...)) {
			c.Next()
			return
		}
		traceId := c.GetHeader(TraceId)
		if traceId == "" {
			traceId = strconv.FormatUint(sonyflakex.NewSonyFlakeId(), 10)
		}

		ctx := logging.NewTraceIDContext(c.Request.Context(), traceId)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set(TraceId, traceId)
		c.Next()
	}
}
