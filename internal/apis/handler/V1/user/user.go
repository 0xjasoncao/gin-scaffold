package user

import (
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (*UserHandler) Login(ctx *gin.Context) {

}

func (*UserHandler) Info(ctx *gin.Context) {

	userInfo := map[string]string{
		"userName": "JasonCao",
		"email":    "codecy2001@gmail.com",
	}

	api.ResSuccess(ctx, userInfo)

}
