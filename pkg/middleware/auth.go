package middleware

import (
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/token"
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
		tokenStr := api.GetToken(c)
		if tokenStr == "" {
			api.ResError(c, errorsx.NewUnauthorized("Unauthorized access."))
			return
		}

		tokenInfo, err := ts.Parse(c, tokenStr)
		if err != nil {
			api.ResError(c, errorsx.NewUnauthorized("Invalid authentication token.").WithError(err))
			return
		}

		ctx := api.ContextWithToken(c, tokenInfo)
		ctx = logging.NewUserIDContext(ctx, tokenInfo.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
