package middleware

import (
	"gin-scaffold/pkg/core/errorsx"
	"gin-scaffold/pkg/core/ginutil"
	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginutil.ResError(c, errorsx.NewMethodNotAllowed("requested method not allowed"))
		return
	}
}
