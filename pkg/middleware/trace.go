package middleware

import (
	"strconv"

	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/0xjasoncao/gin-scaffold/pkg/sonyflakex"
	"github.com/gin-gonic/gin"
)

const TraceId = "X-Request-Id"

func Trace(skipFunc ...SkipFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		if NeedSkip(c, skipFunc...) {
			c.Next()
			return
		}
		traceId := c.GetHeader(TraceId)
		if traceId == "" {
			traceId = strconv.FormatInt(sonyflakex.NewSonyFlakeId(), 10)
		}

		ctx := logging.NewTraceIDContext(c.Request.Context(), traceId)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set(TraceId, traceId)
		c.Next()
	}
}
