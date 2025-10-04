package api

import (
	"strings"

	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/gin-gonic/gin"
)

const (
	ReqBodyKey = "request_body"
)

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

func ParseJSON(c *gin.Context, obj any) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return errors.NewInvalidParams(err)
	}
	return nil
}
