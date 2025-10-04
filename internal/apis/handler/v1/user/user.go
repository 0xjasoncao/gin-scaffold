package user

import (
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) Info(ctx *gin.Context) {

	userInfo := map[string]string{
		"userName": "JasonCao",
		"email":    "codecy2001@gmail.com",
	}

	api.ResData(ctx, userInfo)

}
