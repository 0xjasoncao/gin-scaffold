package user

import (
	"github.com/0xjasoncao/gin-scaffold/internal/apis/request"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/response"
	"github.com/0xjasoncao/gin-scaffold/internal/service"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/token"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/structure"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginSrv service.LoginService
	TokenSrv token.Service
}

func NewLoginHandler(loginSrv service.LoginService, tokenSrv token.Service) *LoginHandler {
	return &LoginHandler{
		LoginSrv: loginSrv,
		TokenSrv: tokenSrv,
	}
}

// Login 登录
//
//	@Summary		login
//	@Description	login by mobile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		request.LoginRequest	true	"登录参数"
//
//	@Success		200		{object}	api.ApiResponse{data=response.LoginResponse}
//	@Router			/login/ [post]
func (h *LoginHandler) Login(ctx *gin.Context) {
	var in request.LoginRequest
	err := api.ParseJSON(ctx, &in)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	user, err := h.LoginSrv.Login(ctx, service.LoginOpt{
		Mobile:   in.Mobile,
		Password: in.Password,
	})
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	res := response.LoginResponse{}
	structure.Copy(&user, &res)

	issuedTokenInfo, err := h.TokenSrv.IssuingToken(ctx, user.ID)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	res.TokenInfo = issuedTokenInfo

	api.ResData(ctx, res)

}

// Logout 登录
//
//	@Summary		logout
//	@Description	login by mobile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"access_token"
//	@Success		200				{object}	api.ApiResponse
//	@Router			/login/exit [post]
func (h *LoginHandler) Logout(ctx *gin.Context) {
	accessToken := api.GetToken(ctx)
	if accessToken == "" {
		api.ResError(ctx, errors.NewUnauthorized("用户未登录"))
		return
	}
	if err := h.TokenSrv.DestroyToken(ctx, accessToken); err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResOK(ctx)

}
