package api

import (
	"github.com/gin-gonic/gin"
	"strings"
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
