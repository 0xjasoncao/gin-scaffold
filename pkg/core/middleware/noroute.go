package middleware

import (
	"gin-scaffold/pkg/core/errorsx"
	"gin-scaffold/pkg/core/ginutil"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginutil.ResError(c, errorsx.NewNotFound("requested URL not found"))
		return
	}
}
