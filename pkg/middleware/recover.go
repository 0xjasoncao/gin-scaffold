package middleware

import (
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logging.WithContext(c.Request.Context()).Error("panic error", zap.Any("panic", err))
				api.ResError(c, errors.NewInternal("Internal server error"))
			}
		}()
		c.Next()
	}
}
