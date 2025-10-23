package middleware

import (
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		api.ResError(c, errorsx.NewNotFound("requested URL not found"))
		return
	}
}
