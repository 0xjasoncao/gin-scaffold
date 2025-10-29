package middleware

import (
	"gin-scaffold/pkg/core/errorsx"
	"gin-scaffold/pkg/core/ginutil"
	"gin-scaffold/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logging.WithContext(c.Request.Context()).Error("panic error", zap.Any("panic", err))
				ginutil.ResError(c, errorsx.NewInternal("Internal server error"))
			}
		}()
		c.Next()
	}
}
