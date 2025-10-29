package middleware

import (
	"gin-scaffold/pkg/core/errorsx"
	ginutil "gin-scaffold/pkg/core/ginutil"
	"gin-scaffold/pkg/core/token"
	"gin-scaffold/pkg/logging"
	"github.com/gin-gonic/gin"
)

type AuthConfig struct {
	Enable            bool     `mapstructure:"enable"`
	SkipPathPrefix    []string `mapstructure:"skip-path-prefix"`
	NotSkipPathPrefix []string `mapstructure:"not-skip-path-prefix"`
}

func Auth(ts token.Service, config *AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if NeedSkip(c, SkippedPathPrefix(config.SkipPathPrefix...), NotSkippedPathPrefix(config.NotSkipPathPrefix...)) {
			c.Next()
			return
		}
		// Request Header: "Authorization": Bearer JWT
		tokenStr := ginutil.GetToken(c)
		if tokenStr == "" {
			ginutil.ResError(c, errorsx.NewUnauthorized("Unauthorized access."))
			return
		}

		tokenInfo, err := ts.Parse(c, tokenStr)
		if err != nil {
			ginutil.ResError(c, errorsx.NewUnauthorized("Invalid authentication token.").WithError(err))
			return
		}

		ctx := ginutil.ContextWithToken(c, tokenInfo)
		ctx = logging.NewUserIDContext(ctx, tokenInfo.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
