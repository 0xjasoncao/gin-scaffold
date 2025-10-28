package v1

import (
	"gin-scaffold/internal/apis/handler/system/request"
	"gin-scaffold/internal/apis/handler/system/response"
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/token"
	"gin-scaffold/pkg/utils/structureutil"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv  system.UserService
	TokenSrv token.Service
}

func NewUserHandler(loginSrv system.UserService, tokenSrv token.Service) *UserHandler {
	user := &UserHandler{
		UserSrv:  loginSrv,
		TokenSrv: tokenSrv,
	}
	return user
}

// Login 登录
//
//	@Summary		login
//	@Description	login
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.LoginRequest	true	"登录参数"
//	@Success		200		{object}	api.Response{data=response.LoginResponse}
//	@Router			/system/user/login/   [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var in request.LoginRequest
	err := api.ParseJSON(ctx, &in)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	user, err := h.UserSrv.Login(ctx.Request.Context(), system.UserQueryParam{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		api.ResError(ctx, err)
		return
	}

	res := response.LoginResponse{}
	structureutil.Copy(&user, &res)

	issuedTokenInfo, err := h.TokenSrv.IssuingToken(ctx, token.Payload{UserID: user.ID})
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	res.TokenInfo = issuedTokenInfo

	api.ResData(ctx, res)

}

// Logout 登出
//
//	@Summary		logout
//	@Description	login by mobile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"access_token"
//	@Success		200				{object}	api.Response
//	@Router			/system/user/logout [post]
func (h *UserHandler) Logout(ctx *gin.Context) {
	accessToken := api.GetToken(ctx)
	if accessToken == "" {
		api.ResError(ctx, errorsx.NewUnauthorized("用户未登录"))
		return
	}
	if err := h.TokenSrv.DestroyToken(ctx.Request.Context(), accessToken); err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResOKWithMessage(ctx, "登出成功")

}
