package api

import (
	"context"
	"gin-scaffold/pkg/token"
	"strings"

	"gin-scaffold/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

type (
	reqTokenInfoKey struct{}
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

// ContextWithToken 将 token.Claims 存入 context
func ContextWithToken(ctx *gin.Context, claims *token.Claims) context.Context {
	return context.WithValue(ctx, reqTokenInfoKey{}, claims)
}

// TokenFromContext 从 context 中获取 token.Claims
func TokenFromContext(ctx context.Context) *token.Claims {
	if v := ctx.Value(reqTokenInfoKey{}); v != nil {
		if claims, ok := v.(*token.Claims); ok {
			return claims
		}
	}
	return nil
}

func ParseJSON(c *gin.Context, obj any) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return errorsx.NewInvalidParams(err)
	}
	return nil
}
