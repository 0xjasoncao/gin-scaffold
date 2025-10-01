package middleware

import (
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/0xjasoncao/gin-scaffold/pkg/token"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
)

func Auth(ts token.Service, skipFunc ...SkipFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if NeedSkip(c, skipFunc...) {
			c.Next()
			return
		}
		tokenStr := api.GetToken(c)
		if tokenStr == "" {
			api.ResError(c, errors.NewUnauthorized("Unauthorized access."))
			return
		}

		tokenInfo, err := ts.Parse(c, tokenStr)
		if err != nil {
			api.ResError(c, errors.NewUnauthorized("Invalid authentication token.").WithError(err))
			return
		}

		ctx := logging.NewUserIDContext(c, tokenInfo.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
