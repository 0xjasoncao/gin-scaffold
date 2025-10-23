package middleware

import (
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		api.ResError(c, errorsx.NewMethodNotAllowed("requested method not allowed"))
		return
	}
}
