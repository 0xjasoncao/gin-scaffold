package middleware

import (
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		api.ResError(c, errors.NewMethodNotAllowed("requested method not allowed"))
		return
	}
}
