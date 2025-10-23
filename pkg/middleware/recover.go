package middleware

import (
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logging.WithContext(c.Request.Context()).Error("panic error", zap.Any("panic", err))
				api.ResError(c, errorsx.NewInternal("Internal server error"))
			}
		}()
		c.Next()
	}
}
