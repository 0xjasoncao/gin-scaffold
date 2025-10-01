package auth

import (
	"github.com/0xjasoncao/gin-scaffold/internal/apis/dto"
	"github.com/0xjasoncao/gin-scaffold/internal/service/auth"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Srv auth.Service
}

func NewHandler(authSrv auth.Service) *Handler {
	return &Handler{Srv: authSrv}
}

// Login 登录
func (auth *Handler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		api.ResError(ctx, errors.NewInvalidParams(err))
		return
	}
	err = auth.Srv.Login(ctx, loginRequest)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResOK(ctx)

}
